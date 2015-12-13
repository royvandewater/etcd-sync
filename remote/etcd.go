package remote

// Etcd defines the interface for communicating with etcd
type Etcd interface {
	List(namespace string) ([]string, error)
}
