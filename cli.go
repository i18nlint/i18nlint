package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "i18nlint"
	app.Usage = "check apps for i18n requirements"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run i18nlint job",
			Action: func(c *cli.Context) error {
				fmt.Println("i18nlint job initiated: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
