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

	client, err := etcd.Dial(etcdURI)
	PanicIfError("etcd.Dial", err)

	keyValues, err := client.KeyValuePairs(namespace)
	PanicIfError("client.KeyValuePairs", err)

	for _, keyValue := range keyValues {
		fmt.Println(keyValue)
	}
}
