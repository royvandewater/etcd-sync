package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/remote"
)

// Dump dumps remote etcd pairs into the local filesystem
func Dump(context *cli.Context) {
	localPath := context.GlobalString("local-path")
	etcdURI := context.GlobalString("etcd-uri")
	remoteDirectory := context.String("remote-directory")

	remoteEtcd, err := remote.Dial(etcdURI)
	PanicIfError("remote.Dial", err)

	keyValues, err := remoteEtcd.KeyValuePairs(remoteDirectory)
	PanicIfError("remote.KeyValuePairs", err)

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
