package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/octoblu/go-simple-etcd-client/etcdclient"
)

// PrintRemote print etcd key/values from remote etcd
func PrintRemote(context *cli.Context) {
	etcdURI := context.GlobalString("etcd-uri")

	etcd, err := etcdclient.Dial(etcdURI)
	if err != nil {
		log.Panicln("etcdclient.Dial failed", err.Error())
	}

	keys, err := etcd.LsRecursive("/")
	if err != nil {
		log.Panicln("etcd.LsRecursive failed", err.Error())
	}

	for _, key := range keys {
		value, err := etcd.Get(key)
		if err != nil {
			log.Panicln("etcd.Get(key)", err.Error())
		}
		if value == "" {
			continue
		}
		fmt.Printf("%v\t%v\n", key, value)
	}
}
