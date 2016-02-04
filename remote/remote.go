package remote

import (
	"github.com/octoblu/go-simple-etcd-client/etcdclient"
	"github.com/royvandewater/etcdsync/keyvalue"
)

// Remote implements etcd and represents the data in a remote etcd server
type Remote struct {
	etcd etcdclient.EtcdClient
}

// Dial creates a Remote from the remote etcd server.
func Dial(etcdURI string) (*Remote, error) {
	etcd, err := etcdclient.Dial(etcdURI)
	return &Remote{etcd}, err
}

// KeyValuePairs returns a list key value pairs
// recursively under the directory
func (remote *Remote) KeyValuePairs(directory string) ([]keyvalue.KeyValue, error) {
	var keyValues []keyvalue.KeyValue
	etcd := remote.etcd

	keys, err := etcd.LsRecursive(directory)
	if err != nil {
		return keyValues, err
	}

	for _, key := range keys {
		value, err := etcd.Get(key)
		if err != nil {
			return make([]keyvalue.KeyValue, 0), err
		}
		if value == "" {
			continue
		}
		keyValues = append(keyValues, keyvalue.KeyValue{key, value})
	}
	return keyValues, nil
}
