# 🍖 gigot

> Gigot is a famous french meal, usually a mutton/lamb leg served with vegetables.                                   

Gigot is a golang-powered secret finder for git repositories, written with performance in mind.

## Example

The simplest way to use gigot is the following:

```go
package main

import (
	"log"
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
	var url = "https://github.com/user/repo"

	// Clone repository
	repo, tempDir, err := gigot.Clone(url)

    // Important: cleanup temporary directory
	defer os.RemoveAll(tempDir)
	if err != nil {
		log.Fatal(err)
	}

	// Init gigot client
	c := gigot.NewClient(signatures, repo)

	log.Print("Starting analysis")
	c.Find(1000)
	log.Print("Analysis has ended")

    // Display findings URLs
	for _, f := range c.Findings {
		log.Printf("%s - %s", f.Match, f.URL(url))
	}
	log.Printf("Found %d matches", len(c.Findings))
}
```
