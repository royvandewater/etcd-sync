package local_test

import (
	"errors"

	"github.com/royvandewater/etcdsync/local"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("local", func() {
	var sut *local.Local
	var dependencies *local.Dependencies
	var mockFS *MockFS
	var err error

	BeforeEach(func() {
		mockFS = &MockFS{}
		dependencies = local.NewDependencies(mockFS)
	})

	Describe("New", func() {
		BeforeEach(func() {
			sut = local.New("the-path", dependencies)
		})

		It("should generate a local instance", func() {
			Expect(sut).NotTo(BeNil())
		})

		It("should set the path", func() {
			Expect(sut.Path).To(Equal("the-path"))
		})
	})

	Describe("Services", func() {
		BeforeEach(func() {
			sut = local.New("the-path", dependencies)
		})

		Context("when there is no directory at the-path", func() {
			BeforeEach(func() {
				mockFS.ReadDirError = errors.New("uh oh")
				_, err = sut.Services()
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})

		Context("when there is a directory at the-path with 1 file", func() {
			var result []local.Service

			BeforeEach(func() {
				fileInfo := &MockFileInfo{NameValue: "the-file"}
				mockFS.ReadDirValue = []local.FileInfo{fileInfo}
				result, err = sut.Services()
			})

			It("should return no error", func() {
				Expect(err).To(BeNil())
			})

			It("should return 1 thing", func() {
				Expect(result).To(HaveLen(1))
			})

			It("should have a path", func() {
				service := result[0]
				Expect(service).NotTo(BeNil())
				Expect(service.Path()).To(Equal("the-path/the-file"))
			})
		})
	})
})
