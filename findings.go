package gigot

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/object"
)

// A Finding is a file line that matched against a signature.
type Finding struct {
	// Signature is the signature the file has matched against.
	Signature Signature
	// Path is the file path relative to the repository root.
	Path string
	// Line is the line number the signature where the signature matched.
	Line int
	// CommitHash is the commit hash in which the finding has been found.
	CommitHash string
	// CommitMsg is the commit message in which the finding has been found.
	CommitMsg string
	// CommitMsg is the commit author in which the finding has been found.
	CommitAuthor string
	// Match is the content that matched against the signature.
	Match string
}

// NewFinding initializes a Finding.
func NewFinding(file *object.File, commit *object.Commit, signature Signature, line int, match string) *Finding {
	return &Finding{
		Signature:    signature,
		Path:         file.Name,
		Line:         line,
		CommitHash:   commit.Hash.String(),
		CommitMsg:    commit.Message,
		CommitAuthor: commit.Author.String(),
		Match:        match,
	}
}

// URL returns the remote URL where the finding can be seen.
func (f *Finding) URL(baseURL string) string {
	return fmt.Sprintf("%s/blob/%s/%s#L%d", baseURL, f.CommitHash, f.Path, f.Line)
}

// Key creates a key from a Finding commit hash, file path, line number and matched content.
func (f *Finding) Key() string {
	return fmt.Sprintf("%s:%s:%d:%s", f.CommitHash, f.Path, f.Line, f.Match)
}
