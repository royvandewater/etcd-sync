package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/local"
)

// PrintLocal prints etcd key/values from local fs
func PrintLocal(context *cli.Context) {
	localPath := context.GlobalString("local-path")
	if !pathIsDir(localPath) {
		ExitWithHelp(fmt.Sprintf("Could not find directory: %v", localPath))
	}

	_, err := local.GenerateHeirarchy(localPath)
	PanicIfError("local.GenerateHierarchy", err)
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
