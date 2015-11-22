package local_test

import "github.com/royvandewater/etcdsync/local"

type MockFS struct {
	ReadDirValue []local.FileInfo
	ReadDirError error

	ReadFileValue []byte
}

func (mockFS *MockFS) ReadDir(dirname string) ([]local.FileInfo, error) {
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
