package remote

// Service implements an Etcd.Service and represents a collection
// of key/value pairs
type Service interface {
	Name() string
}

type localService struct{}

// NewService generates a new Service instance
func NewService() Service {
	return &localService{}
}

func (service *localService) Name() string {
	return ""
}
