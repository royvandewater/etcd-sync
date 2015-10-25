package etcdfs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEtcdfs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Etcdfs Suite")
}
