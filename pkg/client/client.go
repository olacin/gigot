package client

import (
	"sync"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	g "github.com/olacin/gigot/pkg/git"
	s "github.com/olacin/gigot/pkg/signatures"
)

type Client struct {
	Signatures []s.Signature
	Repository *git.Repository
	mutex      sync.Mutex
	Findings   []s.Finding
	seen       map[string]bool
}

func New(signatures []s.Signature, repository *git.Repository) *Client {
	findings := make([]s.Finding, 0)
	seen := make(map[string]bool)
	return &Client{
		Signatures: signatures,
		Repository: repository,
		Findings:   findings,
		seen:       seen,
	}
}

func (c *Client) find(commits <-chan *object.Commit, findings chan<- s.Finding, wg *sync.WaitGroup) {
	defer wg.Done()

	for commit := range commits {
		for _, sig := range c.Signatures {
			matches, err := sig.Find(commit)
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

func (c *Client) Find(workers int) error {
	commits, err := g.Commits(c.Repository)
	if err != nil {
		return err
	}

	jobs := make(chan *object.Commit)
	findings := make(chan s.Finding)

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
