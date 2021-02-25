package npm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type SearchOptions struct {
	Size        uint8   `json:"size"`
	From        uint8   `json:"from"`
	Quality     float32 `json:"quality"`
	Popularity  float32 `json:"popularity"`
	Maintenance float32 `json:"maintenance"`
}

type PackageLinks struct {
	Npm        string `json:"npm"`
	Homepage   string `json:"homepage"`
	Repository string `json:"respository"`
	Bugs       string `json:"bugs"`
}

type PackageAuthor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	URL      string `json:"url"`
	Username string `json:"username"`
}

type PackagePublisher struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PackageFlags struct {
	Unstable bool `json:"unstable"`
}

type PackageScoreDetail struct {
	Quality     float32 `json:"quality"`
	Popularity  float32 `json:"popularity"`
	Maintenance float32 `json:"maintenance"`
}

type PackageScore struct {
	Final       float32            `json:"final"`
	Detail      PackageScoreDetail `json:"detail"`
	SearchScore float32            `json:"searchScore"`
}

type Package struct {
	Name        string             `json:"name"`
	Scope       string             `json:"scope"`
	Version     string             `json:"version"`
	Description string             `json:"description"`
	Keywords    []string           `json:"keywords"`
	Date        string             `json:"date"`
	Links       PackageLinks       `json:"links"`
	Author      PackageAuthor      `json:"author"`
	Publisher   PackagePublisher   `json:"publisher"`
	Maintainers []PackagePublisher `json:"maintainers"`
}

type PackageItem struct {
	Package Package      `json:"package"`
	Flags   PackageFlags `json:"flags"`
	Score   PackageScore `json:"score"`
}

type SearchResults struct {
	Objects []PackageItem `json:"objects"`
	Total   uint32        `json:"total"`
	Time    string        `json:"time"`
}

// Search searches the NPM registry for packages matching the corresponding "pkg" text
func Search(pkg string, opts SearchOptions) (SearchResults, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/-/v1/search?text=%s", pkg)
	if opts.Size != 0 {
		url += "&size=" + string(opts.Size)
	}
	if opts.From != 0 {
		url += "&from=" + string(opts.From)
	}
	if opts.Quality != 0 {
		url += "&quality=" + fmt.Sprintf("%.6f", opts.Quality)
	}
	if opts.Popularity != 0 {
		url += "&popularity=" + fmt.Sprintf("%.6f", opts.Popularity)
	}
	if opts.Maintenance != 0 {
		url += "&maintenance=" + fmt.Sprintf("%.6f", opts.Maintenance)
	}

	resp, err := http.Get(url)
	if err != nil {
		return SearchResults{}, errors.New(fmt.Sprintf("failed to fetch package: %s", err))
	} else {
		var r SearchResults
		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&r)

		return r, nil
	}
}
