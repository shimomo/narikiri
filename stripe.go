package main

import (
	"net/url"
	"strings"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhookendpoint"
)

type Stripe struct {
	SecretKey string
}

func (s *Stripe) ReplaceURLs(new string) error {
	stripe.Key = s.SecretKey
	iterator := webhookendpoint.List(&stripe.WebhookEndpointListParams{})
	for iterator.Next() {
		webhook := iterator.WebhookEndpoint()
		parse, _ := url.Parse(webhook.URL)
		params := &stripe.WebhookEndpointParams{
			URL: stripe.String(strings.Replace(webhook.URL, parse.Scheme+"://"+parse.Host, new, 1)),
		}
		if _, err := webhookendpoint.Update(webhook.ID, params); err != nil {
			return err
		}
	}
	return nil
}
