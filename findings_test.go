package gigot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* func TestNewFinding(t *testing.T) {
    NewFinding()
} */

func TestURL(t *testing.T) {
	baseURL := "https://github.com/user/repo"

	cases := map[string]Finding{
		fmt.Sprintf("%s/blob/1e36bfe10404cb77c12f6dfc8665564f3a41ad7e/README.md#L33", baseURL): {
			CommitHash: "1e36bfe10404cb77c12f6dfc8665564f3a41ad7e",
			Path:       "README.md",
			Line:       33,
		},
		fmt.Sprintf("%s/blob/15686f366981015a7087cbc0ce80df206d8d4150/enum/io.go#L786", baseURL): {
			CommitHash: "15686f366981015a7087cbc0ce80df206d8d4150",
			Path:       "enum/io.go",
			Line:       786,
		},
		fmt.Sprintf("%s/blob/816d787157ccadf5fd92d6b5a0d5b2885dcd7345/resolvers/subdirectory/wildcards.go#L12", baseURL): {
			CommitHash: "816d787157ccadf5fd92d6b5a0d5b2885dcd7345",
			Path:       "resolvers/subdirectory/wildcards.go",
			Line:       12,
		},
	}

	for url, f := range cases {
		assert.Equal(t, f.URL(baseURL), url)
	}
}

func TestKey(t *testing.T) {
	cases := map[string]Finding{
		"1e36bfe10404cb77c12f6dfc8665564f3a41ad7e:README.md:33:password=": {
			CommitHash: "1e36bfe10404cb77c12f6dfc8665564f3a41ad7e",
			Path:       "README.md",
			Line:       33,
			Match:      "password=",
		},
		"15686f366981015a7087cbc0ce80df206d8d4150:enum/io.go:786:apiKey": {
			CommitHash: "15686f366981015a7087cbc0ce80df206d8d4150",
			Path:       "enum/io.go",
			Line:       786,
			Match:      "apiKey",
		},
		"816d787157ccadf5fd92d6b5a0d5b2885dcd7345:resolvers/subdirectory/wildcards.go:12:BEGIN PRIVATE KEY": {
			CommitHash: "816d787157ccadf5fd92d6b5a0d5b2885dcd7345",
			Path:       "resolvers/subdirectory/wildcards.go",
			Line:       12,
			Match:      "BEGIN PRIVATE KEY",
		},
	}

	for url, f := range cases {
		assert.Equal(t, f.Key(), url)
	}
}
