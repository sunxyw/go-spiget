package example

import (
	"context"
	"fmt"

	"github.com/sunxyw/go-spiget/spiget"
)

func RegisterWebhook() {
	client := spiget.NewClient(nil)

	webhook, _, err := client.Webhook.Register(context.Background(), "https://example.com/webhook", []string{"resource-update"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", webhook)
	// &{ID:2a34ebef22f24ff4d... Secret:d8740a0663a193d4...}
}
