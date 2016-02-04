package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/fs"
)

// PrintFS prints etcd key/values from local fs
func PrintFS(context *cli.Context) {
	namespace := context.GlobalString("namespace")
	localPath := context.GlobalString("local-path")
	localEtcdFS := fs.New(localPath)

	keyValues, err := localEtcdFS.KeyValuePairs(namespace)
	PanicIfError("localEtcdFS.KeyValuePairs", err)

	for _, keyValue := range keyValues {
		fmt.Println(keyValue)
	}
}

func pathIsDir(path string) bool {
	if path == "" {
		return false
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
