package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/etcd"
	"github.com/royvandewater/etcdsync/fs"
)

// Load dumps remote etcd pairs into the local filesystem
func Load(context *cli.Context) {
	localPath := context.GlobalString("local-path")
	etcdURI := context.GlobalString("etcd-uri")
	namespace := context.String("namespace")

	localEtcdFS := fs.New(localPath)
	keyValues, err := localEtcdFS.KeyValuePairs(namespace)
	PanicIfError("localEtcdFS.KeyValuePairs", err)

	etcdClient, err := etcd.Dial(etcdURI)
	PanicIfError("etcd.Dial", err)
	err = etcdClient.SetAll(keyValues)
	PanicIfError("etcdClient.SetAll", err)

	os.Exit(0)
}
