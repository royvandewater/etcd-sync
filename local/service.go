package local

import "path"

// Service implements an Etcd.Service and represents a collection
// of key/value pairs
type Service interface {
	Path() string
}

type localService struct {
	dirname  string
	basename string
}

// NewService generates a new Service instance
func NewService(dirname, basename string) Service {
	return &localService{dirname: dirname, basename: basename}
}

func (service *localService) Path() string {
	return path.Join(service.dirname, service.basename)
}
