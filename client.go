package gigot

import (
	"sync"

	"github.com/go-git/go-git/v5/plumbing/object"
)

// Client provides an instance to search Findings in a Git repository.
type Client struct {
	// Signatures is the list of signatures to search for.
	Signatures []Signature
	// Repository is the Git repository to be analyzed.
	Repository Repository
	// mutex is a mutex that prevents goroutines from concurrent accesses.
	mutex sync.Mutex
	// Findings is the list of Findings in a Git repository.
	Findings []Finding
	// seen allows a Client to avoid duplicate findings.
	seen map[string]bool
}

// NewClient initializes a Client.
func NewClient(signatures []Signature, repository Repository) *Client {
	findings := make([]Finding, 0)
	seen := make(map[string]bool)
	return &Client{
		Signatures: signatures,
		Repository: repository,
		Findings:   findings,
		seen:       seen,
	}
}

// find is a goroutine which searches for Findings.
func (c *Client) find(commits <-chan *object.Commit, findings chan<- Finding, wg *sync.WaitGroup) {
	defer wg.Done()

	for commit := range commits {
		for _, sig := range c.Signatures {
			matches, err := sig.Find(c.Repository, commit)
			if err != nil {
				continue
			}

			for _, finding := range matches {
				c.mutex.Lock()
				if _, exists := c.seen[finding.Key()]; exists {
					continue
				}

				c.seen[finding.Key()] = true
				c.mutex.Unlock()

				findings <- finding
			}
		}
	}
}

// Find start a user-defined number of workers to search Findings in a Git repository.
func (c *Client) Find(workers int) error {
	commits, err := c.Repository.Commits()
	if err != nil {
		return err
	}

	jobs := make(chan *object.Commit)
	findings := make(chan Finding)

	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go c.find(jobs, findings, &wg)
	}

	// Send commits into channel
	for _, commit := range commits {
		jobs <- commit
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(findings)
	}()

	// Gather findings in client
	for finding := range findings {
		c.Findings = append(c.Findings, finding)
	}

	return nil
}
