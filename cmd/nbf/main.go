package main

import (
	"log"
	"notion-backup-fetcher/internal/nbf"
	"os"

	"github.com/urfave/cli"
)

func main() {
app := cli.NewApp()
	app.Name = "notion-backup-fetcher"
	app.Usage = "nbf --url https://www.notion.so/kukulam/root-123456"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "url",
			Usage:    "root notion page url e.x. https://www.notion.so/[username]/[page-id]",
			Required: true,
		},
	}
	app.Action = nbf.FetchCommand
	
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
