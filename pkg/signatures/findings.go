package signatures

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type Finding struct {
	Signature    Signature
	Path         string
	Line         int
	CommitHash   string
	CommitMsg    string
	CommitAuthor string
	Match        string
}

func New(file *object.File, commit *object.Commit, signature Signature, line int, match string) *Finding {
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

func (f *Finding) URL(baseURL string) string {
	return fmt.Sprintf("%s/blob/%s/%s#L%d", baseURL, f.CommitHash, f.Path, f.Line)
}
