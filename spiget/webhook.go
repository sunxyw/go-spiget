package spiget

import (
	"context"
	"fmt"
)

type WebhookService service

type Webhook struct {
	ID     string `json:"id,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type WebhookEvents struct {
	Events []string `json:"events"`
}

// Delete a Webhook.
//
// Spiget API docs: https://spiget.org/documentation/#!/webhook/delete_webhook_delete_id_secret
func (w *WebhookService) Delete(ctx context.Context, webhook Webhook) (*Response, error) {
	u := fmt.Sprintf("webhook/delete/%s/%s", webhook.ID, webhook.Secret)
	req, err := w.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return w.client.Do(ctx, req, nil)
}

// Get a list of available events.
//
// Spiget API docs: https://spiget.org/documentation/#!/webhook/get_webhook_events
func (w *WebhookService) GetEvents(ctx context.Context) (*WebhookEvents, *Response, error) {
	u := "webhook/events"
	req, err := w.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var events WebhookEvents
	resp, err := w.client.Do(ctx, req, &events)
	if err != nil {
		return nil, resp, err
	}
	return &events, resp, nil
}

// Register a new Webhook.
//
// Spiget API docs: https://spiget.org/documentation/#!/webhook/post_webhook_register
func (w *WebhookService) Register(ctx context.Context, callbackUrl string, events []string) (*Webhook, *Response, error) {
	u := "webhook/register"
	req, err := w.client.NewRequest("POST", u, map[string]interface{}{
		"url":    callbackUrl,
		"events": events,
	})
	if err != nil {
		return nil, nil, err
	}
	var webhook Webhook
	resp, err := w.client.Do(ctx, req, &webhook)
	if err != nil {
		return nil, resp, err
	}
	return &webhook, resp, nil
}

type WebhookStatus struct {
	Status            int `json:"status"`
	FailedConnections int `json:"failedConnections"`
}

// Get the status of a Webhook.
//
// Spiget API docs: https://spiget.org/documentation/#!/webhook/get_webhook_status_id
func (w *WebhookService) GetStatus(ctx context.Context, id string) (*WebhookStatus, *Response, error) {
	u := "webhook/status/" + id
	req, err := w.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var status WebhookStatus
	resp, err := w.client.Do(ctx, req, &status)
	if err != nil {
		return nil, resp, err
	}
	return &status, resp, nil
}
