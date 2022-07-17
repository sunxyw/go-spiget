package spiget

import "context"

type SearchService service

// Search authors.
//
// Note: This is actually an alias for the AuthorsService.Search method.
func (s *SearchService) SearchAuthor(ctx context.Context, query string, opts *AuthorSearchOptions) ([]*Author, *Response, error) {
	return s.client.Authors.Search(ctx, query, opts)
}

// Search resources.
//
// Note: This is actually an alias for the ResourcesService.Search method.
func (s *SearchService) SearchResource(ctx context.Context, query string, opts *ResourceSearchOptions) ([]*Resource, *Response, error) {
	return s.client.Resources.Search(ctx, query, opts)
}
