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

// ClientDial dials up an etcd client
type ClientDial func(url string) (etcdclient.EtcdClient, error)

// Dial creates a Etcd from the remote etcd server. Leave clientDial
// nil to use the default etcd client
func Dial(etcdURI string, clientDial ClientDial) (*Etcd, error) {
	if clientDial == nil {
		clientDial = etcdclient.Dial
	}
	client, err := clientDial(etcdURI)
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

// Set sets the keyValue on the remote Etcd
func (etcd *Etcd) Set(keyValue keyvalue.KeyValue) error {
	if keyValue.IsDir {
		debug("MkDir: %v", keyValue.Key)
		err := etcd.client.MkDir(keyValue.Key)

		return err
	}

	debug("Set: %v", keyValue)
	return etcd.client.Set(keyValue.Key, keyValue.Value)
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
