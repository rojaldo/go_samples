package main

import (
	"fmt"
	"os"
	"sample/apod"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "date",
			Value: "today",
			Usage: "date to get data for (format: YYYY-MM-DD) (default: today)",
		},
	}
	app.Action = func(c *cli.Context) error {
		if apod.IsValidDate(c.String("date")) {
			fmt.Println("Valid date: ", c.String("date"))
			apod := apod.Apod{}
			apod = apod.GetApod(c.String("date"))
			fmt.Println(apod.Hdurl)
		} else {
			fmt.Println("Invalid date: ", c.String("date"))
		}
		return nil
	}

	app.Run(os.Args)
}
