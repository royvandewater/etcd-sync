package main

import (
	"github.com/codegangsta/cli"
	"github.com/royvandewater/etcdsync/remote"
)

// PrintRemote print etcd key/values from remote etcd
func PrintRemote(context *cli.Context) {
	etcdURI := context.GlobalString("etcd-uri")

	remoteEtcd := remote.New(etcdURI, nil)
	// services, err := remoteEtcd.Services()
	// PanicIfError("remoteEtcd.Services()", err)
	//
	// for _, service := range services {
	// 	fmt.Printf("%v:\n\n", service.Name())
	//
	// 	records, err := service.Records()
	// 	PanicIfError(service.Name(), err)
	// 	for key, value := range records {
	// 		fmt.Printf("%v %v\n", key, value)
	// 	}
	//
	// 	fmt.Print("\n")
	// }
}
