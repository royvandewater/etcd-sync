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
			Name:    "printfs",
			Aliases: []string{"pf"},
			Usage:   "print etcd key/values from local fs",
			Action:  PrintFS,
		},
		{
			Name:    "printetcd",
			Aliases: []string{"pe"},
			Usage:   "print etcd key/values from remote etcd",
			Action:  PrintEtcd,
			Flags:   []cli.Flag{},
		},
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "dump remote etcd pairs into the local filesystem",
			Action:  Dump,
		},
	}
	app.Run(os.Args)
}
