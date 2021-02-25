package npm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SearchOptions options for controlling NPM's response
type SearchOptions struct {
	Size        uint8   `json:"size"`
	From        uint8   `json:"from"`
	Quality     float32 `json:"quality"`
	Popularity  float32 `json:"popularity"`
	Maintenance float32 `json:"maintenance"`
}

// PackageLinks the package's links to other places
type PackageLinks struct {
	Npm        string `json:"npm"`
	Homepage   string `json:"homepage"`
	Repository string `json:"respository"`
	Bugs       string `json:"bugs"`
}

// PackageAuthor information about the package author
type PackageAuthor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	URL      string `json:"url"`
	Username string `json:"username"`
}

// PackagePublisher information about the package's publisher
type PackagePublisher struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PackageUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PackageFlags
type PackageFlags struct {
	Unstable bool `json:"unstable"`
}

// PackageScoreDetail information about the search result ranking for the package
type PackageScoreDetail struct {
	Quality     float32 `json:"quality"`
	Popularity  float32 `json:"popularity"`
	Maintenance float32 `json:"maintenance"`
}

// PackageScore information about the search result rankings for the package
type PackageScore struct {
	Final       float32            `json:"final"`
	Detail      PackageScoreDetail `json:"detail"`
	SearchScore float32            `json:"searchScore"`
}

// Package information about a package
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

// PackageItem a package, its flags, and its ranking scores
type PackageItem struct {
	Package Package      `json:"package"`
	Flags   PackageFlags `json:"flags"`
	Score   PackageScore `json:"score"`
}

// SearchResults overall response returned by the NPM API
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
		return SearchResults{}, fmt.Errorf("failed to fetch package: %s", err)
	}

	if resp.StatusCode != 200 {
		return SearchResults{}, fmt.Errorf("npm responded with something other than 200: %v", resp.StatusCode)
	}

	var r SearchResults
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&r)

	return r, nil
}

type Dist struct {
	Integrity    string `json:"integrity"`
	Shasum       string `json:"shasum"`
	Tarball      string `json:"tarball"`
	FileCount    uint16 `json:"fileCount"`
	UnpackedSize uint32 `json:"unpackedSize"`
	NpmSignature string `json:"npm-signature"`
}

type Bugs struct {
	URL string `json:"url"`
}

type GetPackageVersion struct {
	Name        string               `json:"name"`
	Version     string               `json:"version"`
	Description string               `json:"description"`
	Main        string               `json:"main"`
	Repository  GetPackageRepository `json:"repository"`
	Keywords    []string             `json:"keywords"`
	Author      GetPackageAuthor     `json:"author"`
	License     string               `json:"license"`
	Bugs        Bugs                 `json:"bugs"`
	Homepage    string               `json:"homepage"`
	GitHead     string               `json:"gitHead"`
	Id          string               `json:"_id"`
	NpmVersion  string               `json:"_npmVersion"`
	NodeVersion string               `json:"_nodeVersion"`
	NpmUser     PackageUser          `json:"_npmUser"`
	Dist        Dist                 `json:"dist"`
	Maintainers []PackageUser        `json:"maintainers"`
	// Directories `json:""`
	HasShrinkwrap bool `json:"_hasShrinkwrap"`
}

type GetPackageTime struct {
	Created  string `json:"created"`
	Modified string `json:"modified"`
	string   string
}

type GetPackageRepository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type GetPackageAuthor struct {
	Name string `json:"name"`
}

type GetPackage struct {
	Id             string                       `json:"_id"`
	Rev            string                       `json:"_rev"`
	Name           string                       `json:"name"`
	DistTags       map[string]string            `json:"dist-tags"`
	Versions       map[string]GetPackageVersion `json:"versions"`
	Time           GetPackageTime               `json:"time"`
	Maintainers    []PackageUser                `json:"maintainers"`
	Description    string                       `json:"description"`
	Homepage       string                       `json:"homepage"`
	Keywords       []string                     `json:"keywords"`
	Repository     GetPackageRepository         `json:"repository"`
	Author         GetPackageAuthor             `json:"author"`
	Bugs           Bugs                         `json:"bugs"`
	License        string                       `json:"license"`
	Readme         string                       `json:"readme"`
	ReadmeFilename string                       `json:"readmeFilename"`
}

func FetchPackage(pkg string) (GetPackage, error) {
	resp, err := http.Get("https://registry.npmjs.org/" + pkg)
	if err != nil {
		return GetPackage{}, fmt.Errorf("failed to fetch package: %s", err)
	}

	if resp.StatusCode != 200 {
		return GetPackage{}, fmt.Errorf("npm responded with something other than 200: %v", resp.StatusCode)
	}

	var r GetPackage
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&r)

	return r, nil
}
