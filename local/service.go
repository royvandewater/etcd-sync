package local

import (
	"path"
	"strings"
)

// Service implements an Etcd.Service and represents a collection
// of key/value pairs
type Service interface {
	Name() string
	Path() string
	Records() (map[string]string, error)
}

type localService struct {
	dirname    string
	basename   string
	filesystem FileSystem
}

// NewService generates a new Service instance
func NewService(dirname, basename string, dependencies *Dependencies) Service {
	filesystem := dependencies.GetFileSystem()
	return &localService{dirname: dirname, basename: basename, filesystem: filesystem}
}

func (service *localService) Name() string {
	return service.basename
}

func (service *localService) Path() string {
	return path.Join(service.dirname, service.basename)
}

func (service *localService) Records() (map[string]string, error) {
	contents, err := service.fileContents()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(contents, "\n")

	records := make(map[string]string, len(lines))
	for _, line := range lines {
		key, value := service.parseLine(line)
		records[key] = value
	}
	return records, nil
}

func (service *localService) fileContents() (string, error) {
	contents, err := service.filesystem.ReadFile(service.Path())
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func (service *localService) parseLine(line string) (key, value string) {
	parts := strings.Split(line, " ")
	return parts[0], parts[1]
}
