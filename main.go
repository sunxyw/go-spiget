package main

import (
	"context"

	"github.com/sunxyw/go-spiget/spiget"
)

func main() {
	client := spiget.NewClient(nil)

	authors, _, err := client.Authors.List(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	for _, author := range authors {
		println(author.Name)
	}
}
