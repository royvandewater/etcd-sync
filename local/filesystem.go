package local

import (
	"io/ioutil"
	"os"
)

// FileSystem defines how to access the file system
type FileSystem interface {
	ReadDir(dirname string) ([]FileInfo, error)
}

// File defines how to access a file
type File interface{}

// FileInfo describes a file
type FileInfo interface {
	Name() string
}

// OSFileSystem implements FileSystem using golang os
type OSFileSystem struct{}

// Open opens the named file for reading.
func (OSFileSystem) Open(name string) (File, error) {
	return os.Open(name)
}

// ReadDir reads the directory named by dirname and returns a list of sorted directory entries.
func (osFileSystem *OSFileSystem) ReadDir(dirname string) ([]FileInfo, error) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	return make([]FileInfo, len(fileInfos)), nil
}

// Stat returns the FileInfo structure describing file.
func (OSFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
