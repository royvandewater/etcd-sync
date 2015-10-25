package local_test

import (
	"github.com/royvandewater/etcdsync/local"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("local", func() {
	var err error

	Describe("GenerateHeirarchy", func() {
		Context("When given a path that doesn't exist", func() {

			BeforeEach(func() {
				_, err = local.GenerateHeirarchy("not-extant")
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
