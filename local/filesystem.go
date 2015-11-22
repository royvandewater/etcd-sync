package local

import (
	"io/ioutil"
	"os"
)

// FileSystem defines how to access the file system
type FileSystem interface {
	ReadDir(dirname string) ([]FileInfo, error)
	ReadFile(filename string) ([]byte, error)
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
func (fileSystem *OSFileSystem) ReadDir(dirname string) ([]FileInfo, error) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	osFileInfos := make([]FileInfo, len(fileInfos))
	for i, fileInfo := range fileInfos {
		osFileInfos[i] = fileInfo
	}

	return osFileInfos, nil
}

// ReadFile reads the file named by filename and returns the contents.
func (fileSystem *OSFileSystem) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Stat returns the FileInfo structure describing file.
func (OSFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
