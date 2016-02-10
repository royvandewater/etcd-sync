package main

import (
	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/etcd"
)

// PrintEtcd print etcd key/values from remote etcd
func PrintEtcd(context *cli.Context) {
	namespace := context.GlobalString("namespace")
	etcdURI := context.GlobalString("etcd-uri")
	useTable := context.Bool("table")

	etcdClient, err := etcd.Dial(etcdURI)
	PanicIfError("etcd.Dial", err)

	keyValues, err := etcdClient.KeyValuePairs(namespace)
	PanicIfError("etcdClient.KeyValuePairs", err)

	printKeyValuePairs(useTable, keyValues)
}
