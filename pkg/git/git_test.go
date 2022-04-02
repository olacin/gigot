package git

import (
	"os"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	_, dir, err := Clone("https://github.com/olacin/olacin")
	defer os.RemoveAll(dir)

	assert.Nil(t, err)
	assert.DirExists(t, dir)
}

func TestCommits(t *testing.T) {
	repo, dir, _ := Clone("https://github.com/olacin/olacin")
	defer os.RemoveAll(dir)

	commits, err := Commits(repo)

	assert.Nil(t, err)
	assert.Greater(t, len(commits), 0)
}

func TestFiles(t *testing.T) {
	repo, dir, _ := Clone("https://github.com/olacin/olacin")
	defer os.RemoveAll(dir)

	commit, _ := repo.CommitObject(plumbing.NewHash("3bf80aaf4587881854af6033befe8874ebdabb9c"))
	files, err := Files(commit)

	for _, f := range files {
		contents, _ := f.Contents()
		t.Log(f.Name, f.Size, contents)
	}

	assert.Nil(t, err)
}
