package gigot

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type RepositoryManager interface {
	Clone() (string, error)
	Commits(repo *git.Repository) ([]*object.Commit, error)
	Files(commit *object.Commit) ([]*object.File, error)
}
