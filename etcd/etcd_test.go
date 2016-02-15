package etcd_test

import (
	"fmt"

	"github.com/octoblu/go-simple-etcd-client/etcdclient"
	"github.com/royvandewater/etcdsync/etcd"
	"github.com/royvandewater/etcdsync/keyvalue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Etcd", func() {
	var sut *etcd.Etcd
	var err error
	var fakeClient *FakeClient

	BeforeEach(func() {
		fakeClient = NewFakeClient()
	})

	Describe("Dial", func() {
		Describe("When etcdclient.Dial returns a client", func() {
			BeforeEach(func() {
				sut, err = etcd.Dial("http://somewhere:1111", ThisFakeClient(fakeClient))
			})

			It("Should not return an error", func() {
				Expect(err).To(BeNil())
			})

			It("Should return a client", func() {
				Expect(sut).NotTo(BeNil())
			})
		})
	})

	Describe("With a sut", func() {
		BeforeEach(func() {
			sut, _ = etcd.Dial("http://somewhere:1111", ThisFakeClient(fakeClient))
		})

		Describe("Set", func() {
			Describe("When etcdclient.Set returns nil", func() {
				BeforeEach(func() {
					err = sut.Set(keyvalue.KeyValue{Key: "this", Value: "that"})
				})

				It("should call set with the key and value", func() {
					Expect(fakeClient.SetCalled).To(BeTrue())
					Expect(fakeClient.SetCalledWith).To(Equal([]string{"this", "that"}))
				})

				It("should return nil", func() {
					Expect(err).To(BeNil())
				})
			})

			Describe("When etcdclient.Set returns an error", func() {
				BeforeEach(func() {
					fakeClient.SetReturns = fmt.Errorf("whoops")
					err = sut.Set(keyvalue.KeyValue{Key: "key", Value: "value"})
				})

				It("should call set with the key and value", func() {
					Expect(fakeClient.SetCalled).To(BeTrue())
					Expect(fakeClient.SetCalledWith).To(Equal([]string{"key", "value"}))
				})

				It("should return the error", func() {
					Expect(err).To(MatchError(fmt.Errorf("whoops")))
				})
			})
		})
	})
})

type FakeClient struct {
	SetCalled     bool
	SetCalledWith []string
	SetReturns    error
}

func (client *FakeClient) Del(key string) error                     { return nil }
func (client *FakeClient) DelDir(key string) error                  { return nil }
func (client *FakeClient) Get(key string) (string, error)           { return "", nil }
func (client *FakeClient) Ls(key string) ([]string, error)          { return []string{}, nil }
func (client *FakeClient) LsRecursive(key string) ([]string, error) { return []string{}, nil }
func (client *FakeClient) Set(key, value string) error {
	client.SetCalledWith = []string{key, value}
	client.SetCalled = true
	return client.SetReturns
}

func NewFakeClient() *FakeClient {
	return &FakeClient{}
}

func ThisFakeClient(fakeClient etcdclient.EtcdClient) etcd.ClientDial {
	return func(string) (etcdclient.EtcdClient, error) {
		return fakeClient, nil
	}
}
