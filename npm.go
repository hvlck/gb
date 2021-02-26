package npm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SearchOptions options for controlling NPM's response
// see https://github.com/npm/registry/blob/master/docs/REGISTRY-API.md#get-v1search for more
// documentation for parameters taken from the above source
type SearchOptions struct {
	// the number of package results to return
	Size uint8 `json:"size"`
	// the offset to return results from
	From uint8 `json:"from"`
	// how much of an effect should quality have on search results
	Quality float32 `json:"quality"`
	// how much of an effect should popularity have on search results
	Popularity float32 `json:"popularity"`
	// how much of an effect should maintenance have on search results
	Maintenance float32 `json:"maintenance"`
}

// PackageLinks the package's links to other places
type PackageLinks struct {
	// link to npm package page
	Npm string `json:"npm"`
	// link to (optional) homepage
	Homepage string `json:"homepage"`
	// link to source code repository
	Repository string `json:"respository"`
	// link to issue/bug tracker
	Bugs string `json:"bugs"`
}

// PackageAuthor information about the package author
type PackageAuthor struct {
	// name of the package author
	Name string `json:"name"`
	// package author's email
	Email string `json:"email"`
	// personal site of the author
	URL string `json:"url"`
	// author's npm username
	Username string `json:"username"`
}

// PackagePublisher information about the package's publisher
type PackagePublisher struct {
	// email of the package's publisher
	Email string `json:"email"`
	// username of the package's publisher
	Username string `json:"username"`
}

// PackageUser returns information about an NPM user
type PackageUser struct {
	// name of the npm user
	Name string `json:"name"`
	// email of the NPM user
	Email string `json:"email"`
}

// PackageFlags todo: consult official documentation for this
type PackageFlags struct {
	Unstable bool `json:"unstable"`
}

// PackageScoreDetail information about the search result ranking for the package
type PackageScoreDetail struct {
	// determines how much quality should factor into search rankings
	Quality float32 `json:"quality"`
	// determines how much popularity should factor into search rankings
	Popularity float32 `json:"popularity"`
	// determines how much maintenance should factor into search rankings
	Maintenance float32 `json:"maintenance"`
}

// PackageScore information about the search result rankings for the package
type PackageScore struct {
	// final search ranking score of the package
	Final float32 `json:"final"`
	// package ranking details
	Detail PackageScoreDetail `json:"detail"`
	// todo: find out what this is
	SearchScore float32 `json:"searchScore"`
}

// Package information about a package
type Package struct {
	// name of the package
	Name string `json:"name"`
	// scope of the package
	// see https://docs.npmjs.com/cli/v7/using-npm/scope
	Scope string `json:"scope"`
	// package version
	Version string `json:"version"`
	// package description
	Description string `json:"description"`
	// package keywords
	Keywords []string `json:"keywords"`
	// date the package was last published
	Date string `json:"date"`
	// links for the package
	Links PackageLinks `json:"links"`
	// information about the package author
	Author PackageAuthor `json:"author"`
	// information about the package publisher
	Publisher PackagePublisher `json:"publisher"`
	// information about package maintainers
	Maintainers []PackagePublisher `json:"maintainers"`
}

// PackageItem a package, its flags, and its ranking scores
type PackageItem struct {
	// information about a package
	Package Package `json:"package"`
	// todo: figure out what this is
	Flags PackageFlags `json:"flags"`
	// search ranking scores
	Score PackageScore `json:"score"`
}

// SearchResults overall response returned by the NPM API
type SearchResults struct {
	// list of packages that are similar to the given text
	Objects []PackageItem `json:"objects"`
	// total results
	Total uint32 `json:"total"`
	// time the request was served
	Time string `json:"time"`
}

// Search searches the NPM registry for packages matching the corresponding "pkg" text
// uses NPM's https://registry.npmjs.org/-/v1/search endpoint
// all items in SearchOptions are optional
// see https://github.com/npm/registry/blob/master/docs/REGISTRY-API.md#get-v1search for more information about result customisation
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

// Dist returns information about the package's distribution
type Dist struct {
	// sha512 integrity string
	Integrity string `json:"integrity"`
	Shasum    string `json:"shasum"`
	// link to package's tarball fill
	Tarball string `json:"tarball"`
	// file count of the package
	FileCount uint16 `json:"fileCount"`
	// amount of bytes in the package
	UnpackedSize uint32 `json:"unpackedSize"`
	// PGP signature
	NpmSignature string `json:"npm-signature"`
}

// Bugs returns information for filing bugs
type Bugs struct {
	// bug filing url
	URL string `json:"url"`
}

// GetPackageVersion returns information about a package revision
type GetPackageVersion struct {
	// name of the package
	Name string `json:"name"`
	// package version
	Version string `json:"version"`
	// package description
	Description string `json:"description"`
	// package's main file
	Main string `json:"main"`
	// url of package's repository
	Repository GetPackageRepository `json:"repository"`
	// keywords of the specific version
	Keywords []string `json:"keywords"`
	// package author
	Author PackageAuthor `json:"author"`
	// package's license
	License string `json:"license"`
	// url for filing bug reports
	Bugs Bugs `json:"bugs"`
	// package's homepage
	Homepage string `json:"homepage"`
	// package's git head
	GitHead string `json:"gitHead"`
	// package's internal ID, corresponds to package + version
	Id string `json:"_id"`
	// version of NPM the package was published with
	NpmVersion string `json:"_npmVersion"`
	// version of node the package was published with
	NodeVersion string `json:"_nodeVersion"`
	// information about the publisher of the version
	NpmUser PackageUser `json:"_npmUser"`
	// information about the source's distribution
	Dist Dist `json:"dist"`
	// package maintainers
	Maintainers []PackageUser `json:"maintainers"`
	// Directories `json:""`
	HasShrinkwrap bool `json:"_hasShrinkwrap"`
}

// GetPackageTime returns the date of creation, revision, and times at which versions were published
type GetPackageTime struct {
	// time the package was created
	Created string `json:"created"`
	// last time the package was modified
	Modified string `json:"modified"`
	// version number: publish time
	string string
}

// GetPackageRepository the url and type of repository for the package
type GetPackageRepository struct {
	// type of repository, e.g. "git"
	Type string `json:"type"`
	// url to repository
	URL string `json:"url"`
}

// GetPackage returns information about a npm package
type GetPackage struct {
	// internal id of the package (just the package name)
	Id             string                       `json:"_id"`
	Rev            string                       `json:"_rev"`
	// package name
	Name           string                       `json:"name"`
	DistTags       map[string]string            `json:"dist-tags"`
	// list of previous versions of the package and associated metadata
	Versions       map[string]GetPackageVersion `json:"versions"`
	// list of previous publishing times
	Time           GetPackageTime               `json:"time"`
	// list of package maintainers
	Maintainers    []PackageUser                `json:"maintainers"`
	// package description
	Description    string                       `json:"description"`
	// package homepage
	Homepage       string                       `json:"homepage"`
	// package keywords
	Keywords       []string                     `json:"keywords"`
	// package repository information
	Repository     GetPackageRepository         `json:"repository"`
	// package author information
	Author         PackageAuthor                `json:"author"`
	// package bug report information
	Bugs           Bugs                         `json:"bugs"`
	// package license
	License        string                       `json:"license"`
	// package readme text
	Readme         string                       `json:"readme"`
	// package readme file name
	ReadmeFilename string                       `json:"readmeFilename"`
}

// FetchPackage fetch information about a specific NPM package
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
