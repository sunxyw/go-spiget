package main

import (
	"context"
	"fmt"

	"github.com/sunxyw/go-spiget/spiget"
)

func main() {
	client := spiget.NewClient(nil)

	status, _, err := client.Status.Get(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
}
