package npm

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	resp, err := Search("test", SearchOptions{})
	if err != nil {
		t.Fatalf("failed to fetch package: %s", err)
	}

	_, err = json.Marshal(resp)

	if err != nil {
		t.Fatalf("failed to deserialize json: %s", err)
	}
}

func TestGetPackage(t *testing.T) {
	resp, err := FetchPackage("taita")

	if err != nil {
		t.Fatalf("failed to get package: %s", err)
	}

	if resp.Name != "taita" {
		t.Fatalf("name is invalid: %s", resp.Name)
	}

	if len(resp.DistTags) < 1 {
		t.Fatalf("invalid number of versions: %v", len(resp.Versions))
	}

	if len(resp.Versions) < 2 {
		t.Fatalf("invalid number of versions: %v", len(resp.Versions))
	}

	if len(resp.Maintainers) <= 0 {
		t.Fatalf("invalid number of maintainers: %v", len(resp.Maintainers))
	}

	if resp.Description != "command palette library" {
		t.Fatalf("description invalid: %s", resp.Description)
	}

	if resp.Homepage != "https://github.com/EthanJustice/taita#readme" {
		t.Fatalf("homepage invalid: %s", resp.Homepage)
	}

	if len(resp.Keywords) <= 0 || resp.Keywords[0] != "command-palette" {
		t.Fatalf("keywords invalid")
	}

	if resp.Repository.Type != "git" || resp.Repository.URL != "git+https://github.com/EthanJustice/taita.git" {
		t.Fatalf("repository invalid: %s | %s", resp.Repository.Type, resp.Repository.URL)
	}

	if resp.Author.Name != "Ethan Justice" {
		t.Fatalf("username not correct: %s", resp.Author.Username)
	}

	if resp.Bugs.URL != "https://github.com/EthanJustice/taita/issues" {
		t.Fatalf("invalid bug url: %s", resp.Bugs.URL)
	}

	if resp.License != "MIT" {
		t.Fatalf("invalid license: %s", resp.License)
	}

	if strings.HasPrefix(resp.Readme, "# taita") != true {
		t.Fatalf("readme not valid, may have been changed")
	}

	if resp.ReadmeFilename != "README.md" {
		t.Fatalf("invalid readme file name: %s", resp.ReadmeFilename)
	}

	_, err = json.Marshal(resp)

	if err != nil {
		t.Fatalf("failed to deserialize json: %s", err)
	}
}
