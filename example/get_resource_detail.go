package example

import (
	"context"
	"fmt"

	"github.com/sunxyw/go-spiget/spiget"
)

func GetResourceDetail() {
	client := spiget.NewClient(nil)

	resource, _, err := client.Resources.Get(context.Background(), 6245) // Get resource with ID 6245, PlaceholderAPI
	if err != nil {
		panic(err)
	}

	fmt.Println(resource)
}
