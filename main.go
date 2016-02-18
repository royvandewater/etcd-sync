package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "etcdsync"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "etcd-uri, e",
			Value:  "http://localhost:2379",
			Usage:  "uri where etcd can be found",
			EnvVar: "ETCDSYNC_ETCD_URI",
		},
		cli.StringFlag{
			Name:   "local-path, l",
			Value:  ".",
			Usage:  "path where etcd file values are stored",
			EnvVar: "ETCDSYNC_LOCAL_PATH",
		},
		cli.StringFlag{
			Name:   "namespace, n",
			Value:  "/",
			Usage:  "the etcd namespace",
			EnvVar: "ETCDSYNC_NAMESPACE",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "dump remote etcd pairs into the local filesystem",
			Action:  Dump,
		},
		{
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "load local fs etcd pairs into remote etcd",
			Action:  Load,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:   "include-directories, i",
					Usage:  "Make etcd directories for empty directories in the fs",
					EnvVar: "ETCDSYNC_INCLUDE_DIRECTORIES",
				},
			},
		},
		{
			Name:    "printfs",
			Aliases: []string{"pf"},
			Usage:   "print etcd key/values from local fs",
			Action:  PrintFS,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:   "table, t",
					Usage:  "Show using a pretty ascii table",
					EnvVar: "ETCDSYNC_TABLE",
				},
				cli.BoolFlag{
					Name:   "include-directories, i",
					Usage:  "Include directories in the local fs",
					EnvVar: "ETCDSYNC_INCLUDE_DIRECTORIES",
				},
			},
		},
		{
			Name:    "printetcd",
			Aliases: []string{"pe"},
			Usage:   "print etcd key/values from remote etcd",
			Action:  PrintEtcd,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:   "table, t",
					Usage:  "Show using a pretty ascii table",
					EnvVar: "ETCDSYNC_TABLE",
				},
			},
		},
	}
	app.Run(os.Args)
}
