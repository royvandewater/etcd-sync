package local

// Local implements etcd and represents the data on the file system
type Local struct {
	fileSystem FileSystem
	Path       string
}

// New creates a Local from the local etcd filesystem
func New(path string, dependencies *Dependencies) *Local {
	fs := dependencies.GetFileSystem()
	return &Local{fileSystem: fs, Path: path}
}

// Services returns a list of etcd services
func (local *Local) Services() ([]Service, error) {
	fileInfos, err := local.fileSystem.ReadDir(local.Path)
	if err != nil {
		return nil, err
	}

	services := make([]Service, len(fileInfos))
	for i, fileInfo := range fileInfos {
		services[i] = NewService(local.Path, fileInfo.Name(), nil)
	}

	return services, err
}
