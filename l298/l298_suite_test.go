package l298_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestL298(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "L298 Suite")
}
