package local

import (
	"errors"

	"github.com/royvandewater/etcdsync/etcd"
)

// FromPath generates a Local from the local etcd filesystem
func FromPath(path string, dependencies ...interface{}) (*etcd.Etcd, error) {
	fs := getDependencies(dependencies)
	return nil, errors.New("uh oh")
}

func getDependencies(dependencies ...interface{}) FileSystem {
	return nil
}
