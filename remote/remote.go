package remote

// Remote implements etcd and represents the data in a remote etcd server
type Remote struct {
	etcd      Etcd
	namespace string
	uri       string
}

// New creates a Remote from the remote etcd server.
// The uri parameter is ignored if an Etcd instance is
// injected in through the dependencies
func New(uri, namespace string, deps *Dependencies) *Remote {
	etcd := deps.GetEtcd(uri)
	return &Remote{etcd, namespace, uri}
}

// Namespace returns the etcd directory namespace
// this Remote was constructed with
func (remote *Remote) Namespace() string {
	return remote.namespace
}

// URI returns the etcd service uri this Remote was
// constructed with
func (remote *Remote) URI() string {
	return remote.uri
}

// Services returns a list of etcd services
func (remote *Remote) Services() ([]Service, error) {
	_, err := remote.etcd.List(remote.namespace)
	if err != nil {
		return nil, err
	}

	services := []Service{}

	return services, nil
}
