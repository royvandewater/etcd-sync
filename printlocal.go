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

	localEtcd := local.New(localPath, nil)
	services, err := localEtcd.Services()
	PanicIfError("localEtcd.Services()", err)

	for _, service := range services {
		fmt.Printf("%v:\n\n", service.Name())

		records, err := service.Records()
		PanicIfError(service.Name(), err)
		for key, value := range records {
			fmt.Printf("%v %v\n", key, value)
		}

		fmt.Print("\n")
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
