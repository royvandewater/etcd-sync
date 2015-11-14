package main

import (
	"fmt"
	"log"
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

	localEtcd := local.New(localPath, nil)
	log.Printf("path: %v", localEtcd.Path)
	services, err := localEtcd.Services()
	PanicIfError("localEtcd.Services", err)

	for _, service := range services {
		log.Printf("service: %v", service.Path())
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
