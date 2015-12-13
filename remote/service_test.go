package remote_test

import (
	"github.com/royvandewater/etcdsync/remote"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	var sut remote.Service

	Describe("Name", func() {
		BeforeEach(func() {
			sut = remote.NewService()
		})

		It("should have a name", func() {
			Expect(sut.Name()).NotTo(BeNil())
		})
	})

})
