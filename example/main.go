package main

import (
	"log"
	"net/url"
	"os"
	"regexp"

	"github.com/olacin/gigot"
)

// Setup your signatures
var signatures = []gigot.Signature{{
	Name:        "Testify",
	Description: "Checks if testify is a project dependency",
	Pattern:     regexp.MustCompile("(?i)testify"),
	Score:       3,
}}

func main() {
	// This can also be a file URI
	u, err := url.Parse("https://github.com/user/repo")
	if err != nil {
		log.Fatal(err)
	}

	// Clone repository
	repo := gigot.Repository{
		URL: u,
	}
	tempDir, err := repo.Clone()
	if err != nil {
		log.Fatal(err)
	}

	// Important - cleanup temporary directory
	defer os.RemoveAll(tempDir)

	// Init gigot client
	// c := gigot.NewClient(signatures, repo)
	//
	// log.Print("Starting analysis")
	// c.Find(1000)
	// log.Print("Analysis has ended")

	// Display findings URLs
	// for _, f := range c.Findings {
	// 	log.Printf("%s - %s", f.Match, f.URL(url))
	// }
	// log.Printf("Found %d matches", len(c.Findings))
}
