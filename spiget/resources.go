package spiget

import (
	"context"
	"strconv"
	"strings"
)

// ResourcesService handles communication with the resource related
// methods of the Spiget API.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources
type ResourcesService service

// Resource represents a resource.
type Resource struct {
	ID             int               `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Tag            string            `json:"tag,omitempty"`
	Contributors   string            `json:"contributors,omitempty"`
	Likes          int               `json:"likes,omitempty"`
	File           File              `json:"file,omitempty"`
	TestedVersions []string          `json:"testedVersions,omitempty"`
	Links          map[string]string `json:"links,omitempty"`
	Rating         Rating            `json:"rating,omitempty"`
	ReleaseDate    int               `json:"releaseDate,omitempty"`
	UpdateDate     int               `json:"updateDate,omitempty"`
	Downloads      int               `json:"downloads,omitempty"`
	External       bool              `json:"external,omitempty"`
	Icon           Icon              `json:"icon,omitempty"`
	Premium        bool              `json:"premium,omitempty"`
	Price          float64           `json:"price,omitempty"`
	Currency       string            `json:"currency,omitempty"`
	Author         Author            `json:"author,omitempty"`
	Category       Category          `json:"category,omitempty"`
	Version        Version           `json:"version,omitempty"`
	Versions       []Version         `json:"versions,omitempty"`
	Updates        []Update          `json:"updates,omitempty"`
	SourceCodeLink string            `json:"sourceCodeLink,omitempty"`
	DonationLink   string            `json:"donationLink,omitempty"`
}

func (r *Resource) String() string {
	return Stringify(r)
}

// ResourceListOptions specifies the optional parameters to the
// ResourcesService.List method.
type ResourceListOptions struct {
	ListOptions
}

func (r *ResourcesService) internalList(ctx context.Context, suffix string, opts *ResourceListOptions) ([]*Resource, *Response, error) {
	u := "resources/" + suffix
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var resources []*Resource
	resp, err := r.client.Do(ctx, req, &resources)
	if err != nil {
		return nil, resp, err
	}

	return resources, resp, nil
}

// Get a list of available resources (premium and free).
func (r *ResourcesService) List(ctx context.Context, opts *ResourceListOptions) ([]*Resource, *Response, error) {
	return r.internalList(ctx, "", opts)
}

// ResourceListOptions specifies the optional parameters to the
// ResourcesService.ListByVersions method.
type ResourceListByVersionsOptions struct {
	Method string `url:"method,omitempty"`

	ListOptions
}

// Get resources for the specified version(s).
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_for_version
func (r *ResourcesService) ListByVersions(ctx context.Context, versions []string, opts ResourceListByVersionsOptions) ([]*Resource, *Response, error) {
	versionsStr := strings.Join(versions, ",")
	u := "resources/for/" + versionsStr
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var resources []*Resource
	resp, err := r.client.Do(ctx, req, &resources)
	if err != nil {
		return nil, resp, err
	}

	return resources, resp, nil
}

// Get a list of available free resources.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_free
func (r *ResourcesService) ListFree(ctx context.Context, opts *ResourceListOptions) ([]*Resource, *Response, error) {
	return r.internalList(ctx, "free", opts)
}

// Get all new resources.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_new
func (r *ResourcesService) ListNew(ctx context.Context, opts *ResourceListOptions) ([]*Resource, *Response, error) {
	return r.internalList(ctx, "new", opts)
}

// Get a list of available premium resources.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_premium
func (r *ResourcesService) ListPremium(ctx context.Context, opts *ResourceListOptions) ([]*Resource, *Response, error) {
	return r.internalList(ctx, "premium", opts)
}

// Get a resource by its ID.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource
func (r *ResourcesService) Get(ctx context.Context, id int) (*Resource, *Response, error) {
	u := "resources/" + strconv.Itoa(id)
	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var resource *Resource
	resp, err := r.client.Do(ctx, req, &resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, nil
}

// Get the resource author.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_author
func (r *ResourcesService) GetAuthor(ctx context.Context, id int) (*Author, *Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/author"
	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var author *Author
	resp, err := r.client.Do(ctx, req, &author)
	if err != nil {
		return nil, resp, err
	}

	return author, resp, nil
}

// Download a resource.
//
// This either redirects to spiget's CDN server (cdn.spiget.org) for a direct download of files hosted on
// spigotmc.org or to the URL of externally hosted resources The external field of a resource should be
// checked before downloading, to not receive any unexpected data.
func (r *ResourcesService) Download(ctx context.Context, id int) (*Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/download"
	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return r.client.Do(ctx, req, nil)
}

// Get reviews of a resource.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_reviews
func (r *ResourcesService) GetReviews(ctx context.Context, id int, opts ListOptions) ([]*Review, *Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/reviews"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var reviews []*Review
	resp, err := r.client.Do(ctx, req, &reviews)
	if err != nil {
		return nil, resp, err
	}

	return reviews, resp, nil
}

// Get updates of a resource.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_updates
func (r *ResourcesService) GetUpdates(ctx context.Context, id int, opts ListOptions) ([]*Update, *Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/updates"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var updates []*Update
	resp, err := r.client.Do(ctx, req, &updates)
	if err != nil {
		return nil, resp, err
	}

	return updates, resp, nil
}

// Get the latest resource update
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_updates_latest
func (r *ResourcesService) GetLatestUpdate(ctx context.Context, id int) (*Update, *Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/updates/latest"
	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var update *Update
	resp, err := r.client.Do(ctx, req, &update)
	if err != nil {
		return nil, resp, err
	}

	return update, resp, nil
}

// Get versions of a resource.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_versions
func (r *ResourcesService) GetVersions(ctx context.Context, id int, opts ListOptions) ([]*Version, *Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/versions"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var versions []*Version
	resp, err := r.client.Do(ctx, req, &versions)
	if err != nil {
		return nil, resp, err
	}

	return versions, resp, nil
}

// Get the latest resource version.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_versions_latest
func (r *ResourcesService) GetLatestVersion(ctx context.Context, id int) (*Version, *Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/versions/latest"
	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var version *Version
	resp, err := r.client.Do(ctx, req, &version)
	if err != nil {
		return nil, resp, err
	}

	return version, resp, nil
}

// Get a specific resource version by its ID.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_versions_version
func (r *ResourcesService) GetVersion(ctx context.Context, id int, version int) (*Version, *Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/versions/" + strconv.Itoa(version)
	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var v *Version
	resp, err := r.client.Do(ctx, req, &v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Download a specific resource version.
//
// Note: This only redirects to the stored download location and might not download a file (i.e. for external resources).
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_resources_resource_versions_version_download
func (r *ResourcesService) DownloadVersion(ctx context.Context, id int, version int) (*Response, error) {
	u := "resources/" + strconv.Itoa(id) + "/versions/" + strconv.Itoa(version) + "/download"
	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return r.client.Do(ctx, req, nil)
}

// ResourceSearchOptions specifies the optional parameters to the ResourcesService.Search method.
type ResourceSearchOptions struct {
	Field string `url:"field,omitempty"`

	ListOptions
}

// Search resources.
//
// Spiget API docs: https://spiget.org/documentation/#!/resources/get_search_resources_query
func (r *ResourcesService) Search(ctx context.Context, query string, opts ResourceSearchOptions) ([]*Resource, *Response, error) {
	u := "search/resources/" + query
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := r.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var resources []*Resource
	resp, err := r.client.Do(ctx, req, &resources)
	if err != nil {
		return nil, resp, err
	}

	return resources, resp, nil
}
