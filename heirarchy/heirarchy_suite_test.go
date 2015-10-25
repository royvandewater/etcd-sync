package heirarchy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHeirarchy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heirarchy Suite")
}
