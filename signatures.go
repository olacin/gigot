package gigot

import (
	"log"
	"regexp"

	"github.com/go-git/go-git/v5/plumbing/object"
)

// Signature is a descriptive pattern to match a file against.
type Signature struct {
	// Name is the name of a signature.
	Name string
	// Description is a descriptive text indicating what the signature searches.
	Description string
	// Pattern is a regular expression pattern to match against.
	Pattern *regexp.Regexp
	// Score is a user-defined score indicating the severity of a Finding.
	Score int
}

// Match tries to find Findings on every line of a file.
func (s Signature) Match(file *object.File) ([]Finding, error) {
	content, err := file.Lines()
	if err != nil {
		return nil, err
	}

	findings := make([]Finding, 0)
	for lineNo, line := range content {
		for _, m := range s.Pattern.FindAllString(line, -1) {
			findings = append(findings, Finding{
				Signature: s,
				Path:      file.Name,
				Line:      lineNo + 1,
				Match:     m,
			})
		}
	}

	return findings, nil
}

// Find tries to find Findings on every file of a commit.
func (s Signature) Find(repository Repository, commit *object.Commit) ([]Finding, error) {
	files, err := repository.Files(commit)
	if err != nil {
		return nil, err
	}

	findings := make([]Finding, 0)

	for _, file := range files {
		fileFindings, err := s.Match(file)
		if err != nil {
			// Skip file on error
			log.Print(err)
			continue
		}

		// Enrich findings data with commit information
		for _, finding := range fileFindings {
			finding.CommitHash = commit.Hash.String()
			finding.CommitAuthor = commit.Author.String()
			finding.CommitMsg = commit.Message

			findings = append(findings, finding)
		}
	}

	return findings, nil
}
