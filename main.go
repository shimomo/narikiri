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
	app.Name = "sync"
	app.Usage = "fight the loneliness!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "key, k",
			Usage: "Stripe test secret key.",
		},
	}
	app.Action = func(c *cli.Context) error {
		ngrok := &Ngrok{BaseURL: "http://localhost:4040"}
		url, _ := ngrok.FetchHttpsURL()
		stripe := &Stripe{SecretKey: c.String("key")}
		err := stripe.ReplaceURLs(url)
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
