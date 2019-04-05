package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	godotenv.Load()
	app := cli.NewApp()
	app.Name = "narikiri"
	app.Version = "1.0.0"
	app.Usage = "Stripe and ngrok synchronization command."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "key, k",
			Value: os.Getenv("STRIPE_TEST_SECRET_KEY"),
			Usage: "Stripe test secret key.",
		},
	}
	app.Action = func(c *cli.Context) error {
		ngrok := &Ngrok{BaseURL: "http://localhost:4040"}
		url, e := ngrok.FetchHttpsURL()
		if e != nil {
			return e
		}
		stripe := &Stripe{SecretKey: c.String("key")}
		err := stripe.ReplaceURLs(url)
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
