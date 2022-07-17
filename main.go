package main

import (
	"context"
	"fmt"

	"github.com/sunxyw/go-spiget/spiget"
)

func main() {
	client := spiget.NewClient(nil)

	resource, _, err := client.Resources.Get(context.Background(), 6245)

	if err != nil {
		panic(err)
	}

	fmt.Println(resource)
}
