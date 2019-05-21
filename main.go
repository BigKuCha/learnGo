package main

import (
	"github.com/bigkucha/learnGo/packages"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	packageCommands := []cli.Command{
		{
			Name:   "fmt",
			Action: packages.Fmt,
		},
		{
			Name:   "bufio",
			Action: packages.Bufio,
			Flags:  []cli.Flag{cli.IntFlag{Name: "m", Value: 1}},
		},
		{
			Name:   "channel",
			Action: packages.Channel,
			Flags:  []cli.Flag{cli.IntFlag{Name: "m"}},
		},
	}
	funcsCommands := []cli.Command{}

	app.Commands = append(packageCommands, funcsCommands...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
