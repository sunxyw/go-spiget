package example

import (
	"context"
	"fmt"

	"github.com/sunxyw/go-spiget/spiget"
)

func SearchResources() {
	client := spiget.NewClient(nil)

	opt := &spiget.ResourceSearchOptions{
		Field: "name",
	}
	resources, _, err := client.Resources.Search(context.Background(), "PlaceholderAPI", opt)
	if err != nil {
		panic(err)
	}

	for _, resource := range resources {
		fmt.Println(resource)
	}
}
