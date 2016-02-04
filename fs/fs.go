package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/royvandewater/etcdsync/keyvalue"
)

// FS and represents the data on the file system
type FS struct {
	path string
}

// New creates a FS from the local etcd filesystem
func New(path string) *FS {
	return &FS{path}
}

// KeyValuePairs returns a list key value pairs
// recursively under the namespace
func (fs *FS) KeyValuePairs(namespace string) ([]keyvalue.KeyValue, error) {
	var keyValuePairs []keyvalue.KeyValue

	dir := path.Join(fs.path, namespace)
	err := filepath.Walk(dir, func(keyValuePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		contents, err := ioutil.ReadFile(keyValuePath)
		if err != nil {
			return err
		}

		key, err := filepath.Rel(fs.path, keyValuePath)
		if err != nil {
			return err
		}
		value := strings.TrimSpace(string(contents))
		keyValuePairs = append(keyValuePairs, keyvalue.KeyValue{Key: key, Value: value})

		return nil
	})

	if err != nil {
		return make([]keyvalue.KeyValue, 0), err
	}

	return keyValuePairs, nil
}

// SetAll sets all keyValues on the local fs
func (fs *FS) SetAll(keyValues []keyvalue.KeyValue) error {
	for _, keyValue := range keyValues {
		dir := path.Join(fs.path, path.Dir(keyValue.Key))
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}

		key := path.Join(fs.path, keyValue.Key)
		value := fmt.Sprintln(keyValue.Value)
		err = ioutil.WriteFile(key, []byte(value), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
