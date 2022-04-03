package main

import (
	"log"
	"os"
	"regexp"

	"github.com/olacin/gigot/pkg/client"
	"github.com/olacin/gigot/pkg/git"
	"github.com/olacin/gigot/pkg/signatures"
)

var sigs = []signatures.Signature{{
	Name:        "Testify",
	Description: "Checks if testify is a project dependency",
	Pattern:     regexp.MustCompile("\"github.com/stretchr/testify/assert\""),
	Score:       3,
}}

func main() {
	var url = "https://github.com/olacin/gigot"

	// Clone repository
	log.Printf("Cloning %s...", url)
	repo, tempDir, err := git.Clone(url)
	defer os.RemoveAll(tempDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cloned repository at %s", tempDir)

	// Init client
	c := client.New(sigs, repo)

	log.Print("Starting analysis")
	c.Find()
	log.Print("Analysis has ended")

	for _, f := range c.Findings {
		log.Printf("%s - %s", f.Match, f.URL(url))
	}
	log.Printf("Found %d matches", len(c.Findings))

	log.Printf("Deleting %s", tempDir)
}
