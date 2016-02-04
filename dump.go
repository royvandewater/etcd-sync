package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/etcd"
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

	for _, keyValue := range keyValues {
		dir := path.Join(localPath, path.Dir(keyValue.Key))
		err = os.MkdirAll(dir, 0755)
		PanicIfError(fmt.Sprintf("failed to mkdir '%s'", dir), err)

		key := path.Join(localPath, keyValue.Key)
		value := fmt.Sprintln(keyValue.Value)
		err = ioutil.WriteFile(key, []byte(value), 0755)
		PanicIfError(fmt.Sprintf("failed to write key '%s'", key), err)
	}

	os.Exit(0)
}
