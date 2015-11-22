package local_test

import (
	"os"
	"time"
)

type MockFS struct {
	ReadDirValue []os.FileInfo
	ReadDirError error

	ReadFileValue []byte
}

func (mockFS *MockFS) ReadDir(dirname string) ([]os.FileInfo, error) {
	return mockFS.ReadDirValue, mockFS.ReadDirError
}

func (mockFS *MockFS) ReadFile(filename string) ([]byte, error) {
	return mockFS.ReadFileValue, nil
}

type MockFileInfo struct {
	NameValue string
}

func (fileInfo *MockFileInfo) Name() string {
	return fileInfo.NameValue
}

func (fileInfo *MockFileInfo) Size() int64 {
	return 0
}

func (fileInfo *MockFileInfo) Mode() os.FileMode {
	return 0777
}

func (fileInfo *MockFileInfo) ModTime() time.Time {
	return time.Now()
}

func (fileInfo *MockFileInfo) IsDir() bool {
	return false
}

func (fileInfo *MockFileInfo) Sys() interface{} {
	return nil
}
