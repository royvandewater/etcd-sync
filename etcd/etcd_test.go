package etcd_test

import (
	"github.com/royvandewater/etcdsync/etcd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("etcd", func() {
	Describe("Etcd", func() {
		Context("When instantiated", func() {
			var sut *etcd.Etcd
			BeforeEach(func() {
				sut = etcd.New()
			})

			It("should be a thing", func() {
				Expect(sut).NotTo(BeNil())
			})
		})
	})

})
