package l1474_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"kumarkunalbd/datastructure/l1474"
)

type testCase struct {
	arg1     *l1474.ListNode
	arg2     int
	arg3     int
	expected *l1474.ListNode
}

var _ = Describe("DeleteNNodesEveryM", func() {
	testCases := make([]testCase, 0)

	BeforeEach(func() {
		testCases = append(testCases, testCase{})
	})

})
