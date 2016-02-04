package etcd

import (
	"github.com/octoblu/go-simple-etcd-client/etcdclient"
	"github.com/royvandewater/etcdsync/keyvalue"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("etcdsync:etcd")

// Etcd and represents the data in a remote etcd server
type Etcd struct {
	client etcdclient.EtcdClient
}

// Dial creates a Etcd from the remote etcd server.
func Dial(etcdURI string) (*Etcd, error) {
	client, err := etcdclient.Dial(etcdURI)
	return &Etcd{client}, err
}

// KeyValuePairs returns a list key value pairs
// recursively under the namespace
func (etcd *Etcd) KeyValuePairs(namespace string) ([]keyvalue.KeyValue, error) {
	var keyValues []keyvalue.KeyValue
	client := etcd.client

	keys, err := client.LsRecursive(namespace)
	if err != nil {
		return keyValues, err
	}

	for _, key := range keys {
		value, err := client.Get(key)
		if err != nil {
			return make([]keyvalue.KeyValue, 0), err
		}
		if value == "" {
			continue
		}
		keyValues = append(keyValues, keyvalue.KeyValue{Key: key, Value: value})
	}
	return keyValues, nil
}

// SetAll sets all keyValues on the remote Etcd
func (etcd *Etcd) SetAll(keyValues []keyvalue.KeyValue) error {
	for _, keyValue := range keyValues {
		err := etcd.Set(keyValue)
		if err != nil {
			return err
		}
	}
	return nil
}

// Set sets the keyValue on the remote Etcd
func (etcd *Etcd) Set(keyValue keyvalue.KeyValue) error {
	debug("Set: %v", keyValue)
	err := etcd.client.Set(keyValue.Key, keyValue.Value)
	return err
}
