package spiget

import (
	"context"
	"strconv"
)

// AuthorsService handles communication with the author related
// methods of the Spiget API.
//
// Spiget API docs: https://spiget.org/documentation/#!/authors/get_authors
type AuthorsService service

// Author represents a Spiget author.
type Author struct {
	ID         int               `json:"id,omitempty"`
	Name       string            `json:"name,omitempty"`
	Icon       Icon              `json:"icon,omitempty"`
	Identities map[string]string // may not be present
}

func (a *Author) String() string {
	return Stringify(a)
}

// AuthorListOptions specifies the optional parameters to the
// AuthorsService.List method.
type AuthorListOptions struct {
	ListOptions
}

// Get a list of available authors.
// Note: This only includes members involved with resources, either being their author
// or having reviewed a resource.
//
// Spiget API docs: https://spiget.org/documentation/#!/authors/get_authors
func (a *AuthorsService) List(ctx context.Context, opts *AuthorListOptions) ([]*Author, *Response, error) {

	u := "authors"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := a.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorss []*Author
	resp, err := a.client.Do(ctx, req, &authorss)
	if err != nil {
		return nil, resp, err
	}

	return authorss, resp, nil
}

// Get fetches a author.
//
// Spiget API docs: https://spiget.org/documentation/#!/authors/get_authors_author
func (a *AuthorsService) Get(ctx context.Context, id int) (*Author, *Response, error) {
	u := "authors/" + strconv.Itoa(id)
	req, err := a.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	author := new(Author)
	resp, err := a.client.Do(ctx, req, author)
	if err != nil {
		return nil, resp, err
	}

	return author, resp, nil
}

// AuthorSearchOptions specifies the optional parameters to the
// AuthorsService.Search method.
type AuthorSearchOptions struct {
	Field string `url:"field,omitempty"`

	ListOptions
}

// Search searches for authors by specified field.
//
// Spiget API docs: https://spiget.org/documentation/#!/authors/get_search_authors_query
func (a *AuthorsService) Search(ctx context.Context, query string) ([]*Author, *Response, error) {
	u := "search/authors/" + query
	req, err := a.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorss []*Author
	resp, err := a.client.Do(ctx, req, &authorss)
	if err != nil {
		return nil, resp, err
	}

	return authorss, resp, nil
}
