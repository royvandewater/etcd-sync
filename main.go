package main

import (
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "etcdsync"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "etcd-uri",
			Value: "localhost:2379",
			Usage: "uri where etcd can be found",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "diff",
			Aliases: []string{"d"},
			Usage:   "show differences between local fs & remote etcd",
			Action:  Diff,
		},
	}
}
