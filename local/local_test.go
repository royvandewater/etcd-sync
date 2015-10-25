package local_test

import (
	"github.com/royvandewater/etcdsync/local"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("local", func() {
	var err error

	Describe("FromPath", func() {
		Context("When given a path that doesn't exist", func() {
			BeforeEach(func() {
				_, err = local.FromPath("not-extant")
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})

		Context("When given a path that exists", func() {
			BeforeEach(func() {
				_, err = local.FromPath("extant")
			})

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
