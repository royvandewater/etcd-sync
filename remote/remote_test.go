package remote_test

import (
	"errors"

	"github.com/royvandewater/etcdsync/remote"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("remote", func() {
	var sut *remote.Remote
	var dependencies *remote.Dependencies
	var err error

	Describe("New", func() {
		BeforeEach(func() {
			dependencies = &remote.Dependencies{}
			sut = remote.New("the-uri", "octets", dependencies)
		})

		It("should generate a new remote instance", func() {
			Expect(sut).NotTo(BeNil())
		})

		It("should set the URI", func() {
			Expect(sut.URI()).To(Equal("the-uri"))
		})

		It("should set the Namespace", func() {
			Expect(sut.Namespace()).To(Equal("octets"))
		})
	})

	Describe("Services", func() {
		var result []remote.Service
		var etcd *MockEtcd

		BeforeEach(func() {
			etcd = &MockEtcd{}
			dependencies = remote.NewDependencies(etcd)
			sut = remote.New("the-uri", "octets", dependencies)
		})

		Context("When the server is not reachable", func() {
			BeforeEach(func() {
				etcd.ListError = errors.New("Server not reachable")
				result, err = sut.Services()
			})

			It("should return no result", func() {
				Expect(result).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})

		Context("When the server responds with one directory", func() {
			BeforeEach(func() {
				etcd.ListValue = []string{"some-directory"}
				result, err = sut.Services()
			})

			It("should return a result", func() {
				Expect(result).NotTo(BeNil())
			})

			It("should return no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
