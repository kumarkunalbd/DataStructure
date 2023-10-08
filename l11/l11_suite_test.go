package l11_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestL11(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "L11 Suite")
}
