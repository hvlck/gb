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

	res, err := json.Marshal(resp)

	if err != nil {
		t.Fatalf("failed to deserialize json: %s", err)
	}
}
