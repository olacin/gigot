package gigot

/* import (
	"os"
	"regexp"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	s := Signature{
		Name:        "url",
		Description: "URL",
		Pattern:     regexp.MustCompile("keychron"),
		Score:       3,
	}

	repo, dir, _ := Clone("https://github.com/olacin/olacin")
	defer os.RemoveAll(dir)

	commit, _ := repo.CommitObject(plumbing.NewHash("3bf80aaf4587881854af6033befe8874ebdabb9c"))
	file, _ := commit.File("README.md")

	matches, err := s.Match(file)

	assert.Nil(t, err)
	for _, m := range matches {
		t.Log(file.Name, m.Line, m.Match, m.Signature.Name)
	}
} */
