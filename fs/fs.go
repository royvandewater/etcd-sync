package fs

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/royvandewater/etcdsync/keyvalue"
)

// FS and represents the data on the file system
type FS struct {
	Path string
}

// New creates a FS from the local etcd filesystem
func New(Path string) *FS {
	return &FS{Path}
}

// KeyValuePairs returns a list key value pairs
// recursively under the namespace
func (fs *FS) KeyValuePairs(namespace string) ([]keyvalue.KeyValue, error) {
	var keyValuePairs []keyvalue.KeyValue

	dir := path.Join(fs.Path, namespace)
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

		Key, err := filepath.Rel(fs.Path, keyValuePath)
		if err != nil {
			return err
		}
		Value := strings.TrimSpace(string(contents))
		keyValuePairs = append(keyValuePairs, keyvalue.KeyValue{Key, Value})

		return nil
	})

	if err != nil {
		return make([]keyvalue.KeyValue, 0), err
	}

	return keyValuePairs, nil
}
