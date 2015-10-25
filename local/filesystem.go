package local

import "os"

// FileSystem defines how to access the file system
type FileSystem interface {
}

// File defines how to access a file
type File interface {
}

// OSFileSystem implements FileSystem using golang os
type OSFileSystem struct{}

// Open opens the named file for reading.
func (OSFileSystem) Open(name string) (File, error) {
	return os.Open(name)
}

// Stat returns the FileInfo structure describing file.
func (OSFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
