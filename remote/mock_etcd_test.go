package remote_test

type MockEtcd struct {
	ListError error
	ListValue []string
}

func (etcd *MockEtcd) List(namespace string) ([]string, error) {
	return etcd.ListValue, etcd.ListError
}
