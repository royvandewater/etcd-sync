package local

import (
	"io/ioutil"
	"os"
)

// FileSystem defines how to access the file system
type FileSystem interface {
	// ReadDir reads the directory named by dirname and returns a list of sorted directory entries.
	ReadDir(dirname string) ([]os.FileInfo, error)
	// ReadFile reads the file named by filename and returns the contents.
	ReadFile(filename string) ([]byte, error)
}

// File defines how to access a file
type File interface{}

// NewFileSystem constructs a new FileSystem
func NewFileSystem() FileSystem {
	return &osFileSystem{}
}

// osFileSystem implements FileSystem using golang os
type osFileSystem struct{}

func (fileSystem *osFileSystem) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func (fileSystem *osFileSystem) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
