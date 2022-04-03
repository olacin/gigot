package signatures

import (
	"regexp"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/olacin/gigot/pkg/git"
)

type Signature struct {
	Name        string
	Description string
	Pattern     *regexp.Regexp
	Score       int
}

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

func (s Signature) Find(commit *object.Commit) ([]Finding, error) {
	files, err := git.Files(commit)
	if err != nil {
		return nil, err
	}

	findings := make([]Finding, 0)

	for _, file := range files {
		fileFindings, err := s.Match(file)
		if err != nil {
			// Skip file on error
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
