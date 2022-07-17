package spiget

import (
	"context"
	"strconv"
)

// AuthorsService handles communication with the category related
// methods of the Spiget API.
//
// Spiget API docs: https://spiget.org/documentation/#!/categories/get_categories
type CategoriesService service

// Category represents a Spiget category.
type Category struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Category) String() string {
	return Stringify(c)
}

// CategoryListOptions specifies the optional parameters to the
// CategoriesService.List method.
type CategoryListOptions struct {
	ListOptions
}

// Get a list of categories.
//
// Spiget API docs: https://spiget.org/documentation/#!/categories/get_categories
func (c *CategoriesService) List(ctx context.Context, opts *CategoryListOptions) ([]*Category, *Response, error) {
	u := "categories"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var categories []*Category
	resp, err := c.client.Do(ctx, req, &categories)
	if err != nil {
		return nil, resp, err
	}

	return categories, resp, nil
}

// Get details about a category
//
// Spiget API docs: https://spiget.org/documentation/#!/categories/get_categories_category
func (c *CategoriesService) Get(ctx context.Context, id int) (*Category, *Response, error) {
	u := "categories/" + strconv.Itoa(id)
	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var category *Category
	resp, err := c.client.Do(ctx, req, &category)
	if err != nil {
		return nil, resp, err
	}

	return category, resp, nil
}
