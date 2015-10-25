package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "etcdsync"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "etcd-uri",
			Value:  "localhost:2379",
			Usage:  "uri where etcd can be found",
			EnvVar: "ETCDSYNC_ETCD_URI",
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
	app.Run(os.Args)
}
