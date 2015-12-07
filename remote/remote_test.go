package remote_test

import (
	"github.com/royvandewater/etcdsync/remote"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("remote", func() {
	var sut *remote.Remote
	var err error

	Describe("New", func() {
		BeforeEach(func() {
			sut = remote.New("the-uri")
		})

		It("should generate a new remote instance", func() {
			Expect(sut).NotTo(BeNil())
		})

		It("should set the Uri", func() {
			Expect(sut.URI).To(Equal("the-uri"))
		})
	})

	Describe("Services", func() {
		BeforeEach(func() {
			sut = remote.New("the-uri")
		})

		Context("When the server does not respond", func() {
			BeforeEach(func() {
				_, err = sut.Services()
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
