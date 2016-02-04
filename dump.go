package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/etcd"
	"github.com/royvandewater/etcdsync/fs"
)

// Dump dumps remote etcd pairs into the local filesystem
func Dump(context *cli.Context) {
	localPath := context.GlobalString("local-path")
	etcdURI := context.GlobalString("etcd-uri")
	namespace := context.String("namespace")

	client, err := etcd.Dial(etcdURI)
	PanicIfError("etcd.Dial", err)

	keyValues, err := client.KeyValuePairs(namespace)
	PanicIfError("client.KeyValuePairs", err)

	localEtcdFS := fs.New(localPath)
	err = localEtcdFS.SetAll(keyValues)
	PanicIfError("localEtcdFS.SetAll", err)

	os.Exit(0)
}
