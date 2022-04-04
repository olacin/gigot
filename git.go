package gigot

import (
	"errors"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Repository struct {
	URL *url.URL
	ref *git.Repository
}

func NewRepository(URL string) (*Repository, error) {
	u, err := url.ParseRequestURI(URL)
	log.Print(u, err)
	if err != nil {
		return nil, err
	}
	return &Repository{
		URL: u,
	}, nil
}

// Clone clones a repository into a temporary directory.
func (r Repository) Clone() (string, error) {
	if r.URL == nil {
		return "", errors.New("repository URL is nil")
	}
	// Create temporary directory to clone repository
	dir, err := ioutil.TempDir("", "gigot_")
	if err != nil {
		return "", err
	}

	// Clone repository
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{URL: r.URL.String()})
	if err != nil {
		return dir, err
	}

	r.ref = repo

	return dir, nil
}

// Commits gathers every commit in a repository history.
func (r Repository) Commits() ([]*object.Commit, error) {
	if r.ref == nil {
		return nil, errors.New("repository has not been cloned yet")
	}

	commits := make([]*object.Commit, 0)

	// Grab every commit ever made in repository
	rCommits, err := r.ref.CommitObjects()
	if err != nil {
		return nil, err
	}

	rCommits.ForEach(func(c *object.Commit) error {
		commits = append(commits, c)
		return nil
	})

	return commits, nil
}

// Files lists all affected files in a single commit.
func (r Repository) Files(commit *object.Commit) ([]*object.File, error) {
	files := make([]*object.File, 0)

	// Grab commit files
	cFiles, err := commit.Files()
	if err != nil {
		return nil, err
	}

	// Skip big files from analysis
	cFiles.ForEach(func(f *object.File) error {
		// TODO: skip big files
		files = append(files, f)
		return nil
	})

	return files, nil
}
