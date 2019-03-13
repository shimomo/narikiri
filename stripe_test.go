package main

import (
	"os"
	"testing"
)

func TestReplaceURLs(t *testing.T) {
	stripe := &Stripe{SecretKey: os.Getenv("STRIPE_TEST_SECRET_KEY")}
	err := stripe.ReplaceURLs("https://a6c25881.ngrok.io")
	if err != nil {
		t.Errorf(err.Error())
	}
}
