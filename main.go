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
			Name:   "etcd-uri, e",
			Value:  "localhost:2379",
			Usage:  "uri where etcd can be found",
			EnvVar: "ETCDSYNC_ETCD_URI",
		},
		cli.StringFlag{
			Name:   "local-path, l",
			Value:  ".",
			Usage:  "path where etcd file values are stored",
			EnvVar: "ETCDSYNC_LOCAL_PATH",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "diff",
			Aliases: []string{"d"},
			Usage:   "show differences between local fs & remote etcd",
			Action:  Diff,
		},
		{
			Name:    "printlocal",
			Aliases: []string{"pl"},
			Usage:   "print etcd key/values from local fs",
			Action:  PrintLocal,
		},
		{
			Name:    "printremote",
			Aliases: []string{"pl"},
			Usage:   "print etcd key/values from remote etcd",
			Action:  PrintRemote,
		},
	}
	app.Run(os.Args)
}
