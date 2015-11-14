package local_test

import (
	"github.com/royvandewater/etcdsync/local"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Describe("Path", func() {
		var sut local.Service

		BeforeEach(func() {
			sut = local.NewService("/path/to/", "file")
		})

		It("should have a Path", func() {
			Expect(sut.Path()).To(Equal("/path/to/file"))
		})
	})
})
