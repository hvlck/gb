package npm

import (
	"encoding/json"
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
	resp, err := FetchPackage("test")

	if err != nil {
		t.Fatalf("failed to get package: %s", err)
	}

	_, err = json.Marshal(resp)

	if err != nil {
		t.Fatalf("failed to deserialize json: %s", err)
	}
}
