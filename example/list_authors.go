package example

import (
	"context"
	"fmt"

	"github.com/sunxyw/go-spiget/spiget"
)

func ListAuthors() {
	client := spiget.NewClient(nil)

	opt := &spiget.AuthorListOptions{
		ListOptions: spiget.ListOptions{
			Page: 3,
		},
	}
	authors, _, err := client.Authors.List(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(authors)
}
