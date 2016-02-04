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
			Name:    "printlocal",
			Aliases: []string{"pl"},
			Usage:   "print etcd key/values from local fs",
			Action:  PrintLocal,
		},
		{
			Name:    "printremote",
			Aliases: []string{"pr"},
			Usage:   "print etcd key/values from remote etcd",
			Action:  PrintRemote,
		},
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "dump remote etcd pairs into the local filesystem",
			Action:  Dump,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "remote-directory, r",
					Value:  "/",
					Usage:  "etcd directory to dump to local fs. Directory is still included in path of dumped files",
					EnvVar: "ETCDSYNC_REMOTE_DIRECTORY",
				},
			},
		},
	}
	app.Run(os.Args)
}
