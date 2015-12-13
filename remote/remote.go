package remote

import "errors"

// Remote implements etcd and represents the data in a remote etcd server
type Remote struct {
	URI string
}

// New creates a Remote from the remote etcd server
func New(uri string) *Remote {
	return &Remote{URI: uri}
}

// Services returns a list of etcd services
func (remote *Remote) Services() ([]Service, error) {
	return nil, errors.New("oops")
}
