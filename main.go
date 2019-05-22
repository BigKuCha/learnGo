package main

import (
	"github.com/bigkucha/learnGo/packages"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "method,m",
			Value: 1,
		},
	}
	packageCommands := []cli.Command{
		{
			Name:   "fmt",
			Action: packages.Fmt,
		},
		{
			Name:   "bufio",
			Action: packages.Bufio,
		},
		{
			Name:   "channel",
			Action: packages.Channel,
		},
		{
			Name:   "context",
			Action: packages.Context,
		},
		{
			Name:   "exec",
			Action: packages.Exec,
		},
		{
			Name:   "flag",
			Action: packages.Flag,
		},
	}
	funcsCommands := []cli.Command{}

	app.Commands = append(packageCommands, funcsCommands...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
