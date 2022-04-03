package client

import (
	"fmt"
	"sync"

	"github.com/go-git/go-git/v5"
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

func (c *Client) SaveFinding(finding s.Finding) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Don't add duplicates
	var key string = fmt.Sprintf("%s:%s:%d:%s", finding.CommitHash, finding.Path, finding.Line, finding.Match)
	if _, exists := c.seen[key]; exists {
		return
	}

	c.seen[key] = true
	c.Findings = append(c.Findings, finding)
}

func (c *Client) Find() error {
	var wg sync.WaitGroup

	commits, err := g.Commits(c.Repository)
	if err != nil {
		return err
	}

	for _, commit := range commits {
		wg.Add(1)

		commit := commit

		go func() {
			defer wg.Done()
			for _, sig := range c.Signatures {
				findings, err := sig.Find(commit)
				if err != nil {
					continue
				}

				for _, finding := range findings {
					c.SaveFinding(finding)
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
