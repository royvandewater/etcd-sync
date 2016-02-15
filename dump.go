package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/etcd"
	"github.com/royvandewater/etcdsync/fs"
)

// Dump dumps remote etcd pairs into the local filesystem
func Dump(context *cli.Context) {
	namespace := context.GlobalString("namespace")
	localPath := context.GlobalString("local-path")
	etcdURI := context.GlobalString("etcd-uri")

	etcdClient, err := etcd.Dial(etcdURI, nil)
	FatalIfError("etcd.Dial", err)

	keyValues, err := etcdClient.KeyValuePairs(namespace)
	FatalIfError("etcdClient.KeyValuePairs", err)

	localEtcdFS := fs.New(localPath)
	err = localEtcdFS.SetAll(keyValues)
	FatalIfError("localEtcdFS.SetAll", err)

	os.Exit(0)
}
