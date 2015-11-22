package local_test

import (
	"github.com/royvandewater/etcdsync/local"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	var dependencies *local.Dependencies
	var mockFS *MockFS
	var sut local.Service

	BeforeEach(func() {
		mockFS = &MockFS{}
		dependencies = local.NewDependencies(mockFS)
		sut = local.NewService("/path/to/", "file.service", dependencies)
	})

	Describe("Name", func() {
		It("should have a Name", func() {
			Expect(sut.Name()).To(Equal("file.service"))
		})
	})

	Describe("Path", func() {
		It("should have a Path", func() {
			Expect(sut.Path()).To(Equal("/path/to/file.service"))
		})
	})

	Describe("Records", func() {
		Context("When the service file has one line", func() {
			BeforeEach(func() {
				mockFS.ReadFileValue = []byte("key value")
			})

			It("should have Records", func() {
				records, err := sut.Records()
				Expect(err).To(BeNil())
				Expect(records).To(HaveLen(1))
			})
		})

		Context("When the service file has two lines", func() {
			BeforeEach(func() {
				mockFS.ReadFileValue = []byte("key1 value1\nkey2 value2")
			})

			It("should have Records", func() {
				records, err := sut.Records()
				Expect(err).To(BeNil())
				Expect(records).To(HaveLen(2))
			})
		})
	})
})
