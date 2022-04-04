package gigot

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGit struct {
	mock.Mock
}

func (m *MockGit) PlainClone() {
	return
}

func TestNewRepository(t *testing.T) {
	// Invalid URL
	repo, err := NewRepository("this is not a valid URL")
	assert.Nil(t, repo)
	assert.NotNil(t, err)

	// Valid URL
	repo, err = NewRepository("https://github.com/olacin/olacin")
	assert.NotNil(t, repo)
	assert.Nil(t, err)
}

func TestClone(t *testing.T) {
	repo, _ := NewRepository("https://github.com/olacin/olacin")

	m := new(MockGit)
	m.On("PlainClone", "https://github.com/olacin/olacin").Return(true)
	// repo, err := git.PlainClone(dir, false, &git.CloneOptions{URL: r.URL.String()})

	dir, err := repo.Clone()
	defer os.RemoveAll(dir)

	m.AssertCalled(t, "PlainClone")

	assert.Nil(t, err)
	assert.DirExists(t, dir)
}

/* func TestCommits(t *testing.T) {
	repo, dir, _ := Clone("https://github.com/olacin/olacin")
	defer os.RemoveAll(dir)

	commits, err := Commits(repo)

	assert.Nil(t, err)
	assert.NotEmpty(t, commits)
}

func TestFiles(t *testing.T) {
	repo, dir, _ := Clone("https://github.com/olacin/olacin")
	defer os.RemoveAll(dir)

	commit, _ := repo.CommitObject(plumbing.NewHash("3bf80aaf4587881854af6033befe8874ebdabb9c"))
	files, err := Files(commit)

	assert.Nil(t, err)
	assert.NotEmpty(t, files)
} */
