package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var language string
	var recusive bool
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "select language",
			Destination: &language,
		},
		cli.BoolFlag{
			Name:        "recusive, r",
			Usage:       "recusive for the greeting",
			Destination: &recusive,
		},
	}

	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Println("cmd is ", cmd)
		}
		fmt.Println("recusive is ", recusive)
		fmt.Println("language is ", language)
		return nil
	}
	app.Run(os.Args)
}
