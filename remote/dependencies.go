package remote

// Dependencies define what can be injected into Remote
type Dependencies struct {
	etcd Etcd
}

// NewDependencies constructs a new dependencies instance
func NewDependencies(etcd Etcd) *Dependencies {
	return &Dependencies{etcd}
}

// GetEtcd returns the injected etcd or generates a new
// instance using github.com/coreos/etcd/client and configured
// with the etcd server uri
func (dependencies *Dependencies) GetEtcd(uri string) Etcd {
	return dependencies.etcd
}
