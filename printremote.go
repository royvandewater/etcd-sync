package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/remote"
)

// PrintRemote print etcd key/values from remote etcd
func PrintRemote(context *cli.Context) {
	etcdURI := context.GlobalString("etcd-uri")

	remoteEtcd, err := remote.Dial(etcdURI)
	PanicIfError("remote.Dial", err)

	keyValues, err := remoteEtcd.KeyValuePairs("/octoblu")
	PanicIfError("remoteEtcd.KeyValuePairs", err)

	for _, keyValue := range keyValues {
		fmt.Println(keyValue)
	}
}
