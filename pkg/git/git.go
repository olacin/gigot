package git

import (
	"io/ioutil"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func Clone(url string) (*git.Repository, string, error) {
	// Create temporary directory to clone repository
	dir, err := ioutil.TempDir("", "gigot_")
	if err != nil {
		return nil, "", err
	}

	// Clone repository
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{URL: url})
	if err != nil {
		return nil, dir, err
	}

	return repo, dir, nil
}

func Commits(repo *git.Repository) ([]*object.Commit, error) {
	commits := make([]*object.Commit, 0)

	// Grab every commit ever made in repository
	rCommits, err := repo.CommitObjects()
	if err != nil {
		return nil, err
	}

	rCommits.ForEach(func(c *object.Commit) error {
		commits = append(commits, c)
		return nil
	})

	return commits, nil
}

func Files(commit *object.Commit) ([]*object.File, error) {
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
