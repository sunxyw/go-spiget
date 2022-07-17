package main

import (
	"context"
	"fmt"

	"github.com/sunxyw/go-spiget/spiget"
)

func main() {
	client := spiget.NewClient(nil)

	webhook, _, err := client.Webhook.Register(context.Background(), "https://example.com/webhook", []string{"resource-update"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", webhook)
}
