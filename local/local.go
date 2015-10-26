package local

// Local implements etcd and represents the data on the file system
type Local struct {
	FileSystem FileSystem
	Path       string
}

// New creates a Local from the local etcd filesystem
func New(path string, dependencies *Dependencies) *Local {
	// _, err := fs.ReadDir(path)
	// if err != nil {
	// 	return nil, err
	// }
	fs := dependencies.GetFileSystem()
	return &Local{FileSystem: fs, Path: path}
}

// Services returns a list of etcd services
func (local *Local) Services() (int, error) {
	_, err := local.FileSystem.ReadDir(local.Path)
	return 0, err
}
