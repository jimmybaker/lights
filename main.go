package main

import (
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "setup",
				Aliases: []string{"s"},
				Usage:   "initial setup",
				Action: func(cCtx *cli.Context) error {
					Setup()
					return nil
				},
			},
			{
				Name:  "on",
				Usage: "turn on lights",
				Action: func(cCtx *cli.Context) error {
					nameOfLamp := strings.Join(cCtx.Args().Slice(), " ")
					Toggle(nameOfLamp, true)
					return nil
				},
			},
			{
				Name:  "off",
				Usage: "turn off lights",
				Action: func(cCtx *cli.Context) error {
					nameOfLamp := strings.Join(cCtx.Args().Slice(), " ")
					Toggle(nameOfLamp, false)
					return nil
				},
			},
			{
				Name:    "brighten",
				Usage:   "brighten lights",
				Aliases: []string{"bri", "b"},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "percent",
						Usage: "brightness percentage value to set",
						Value: 100,
					},
				},
				Action: func(cCtx *cli.Context) error {
					nameOfLamp := strings.Join(cCtx.Args().Slice(), " ")
					Brighten(nameOfLamp, cCtx.Int("percent"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
