package heirarchy_test

import (
	"github.com/royvandewater/etcdsync/heirarchy"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("heirarchy", func() {
	Describe("Heirarchy", func() {
		Context("When instantiated", func() {
			var sut *heirarchy.Heirarchy
			BeforeEach(func() {
				sut = heirarchy.New()
			})

			It("should be a thing", func() {
				Expect(sut).NotTo(BeNil())
			})
		})
	})

})
