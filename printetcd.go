package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/etcd"
)

// PrintEtcd print etcd key/values from remote etcd
func PrintEtcd(context *cli.Context) {
	namespace := context.GlobalString("namespace")
	etcdURI := context.GlobalString("etcd-uri")

	etcdClient, err := etcd.Dial(etcdURI)
	PanicIfError("etcd.Dial", err)

	keyValues, err := etcdClient.KeyValuePairs(namespace)
	PanicIfError("etcdClient.KeyValuePairs", err)

	for _, keyValue := range keyValues {
		fmt.Println(keyValue)
	}
}
