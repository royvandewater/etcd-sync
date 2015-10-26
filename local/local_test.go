package local_test

import (
	"errors"

	"github.com/royvandewater/etcdsync/local"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockFS struct {
	ReadDirValue []local.FileInfo
	ReadDirError error
}

func (mockFS *MockFS) ReadDir(dirname string) ([]local.FileInfo, error) {
	return mockFS.ReadDirValue, mockFS.ReadDirError
}

var _ = Describe("local", func() {
	var sut *local.Local
	var dependencies *local.Dependencies
	var mockFS *MockFS

	BeforeEach(func() {
		mockFS = &MockFS{}
		dependencies = &local.Dependencies{FileSystem: mockFS}
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

		It("should set FileSystem", func() {
			Expect(sut.FileSystem).To(Equal(mockFS))
		})
	})

	Describe("Services", func() {
		BeforeEach(func() {
			sut = local.New("the-path", dependencies)
		})

		Context("when there is no directory at the-path", func() {
			var err error

			BeforeEach(func() {
				mockFS.ReadDirError = errors.New("uh oh")
				_, err = sut.Services()
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
