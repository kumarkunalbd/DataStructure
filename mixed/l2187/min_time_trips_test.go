package l2187_test

import (
	"kumarkunalbd/datastructure/l2187"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type testCase struct {
	arg1     []int
	arg2     int
	expected int64
}

var _ = Describe("MinTimeTrips", func() {
	testCases := make([]testCase, 0)
	BeforeEach(func() {
		testCases = append(testCases, testCase{[]int{1, 2, 3}, 5, 3}, testCase{[]int{2}, 1, 2})
	})

	Describe("Minimu time", func() {
		Context("first test case", func() {
			It("Should be matching", func() {
				Expect(l2187.MinimumTime(testCases[0].arg1, int(testCases[0].arg2))).To(Equal(testCases[0].expected))
			})
		})

		Context("second test case", func() {
			It("Should be matching", func() {
				Expect(l2187.MinimumTime(testCases[1].arg1, int(testCases[1].arg2))).To(Equal(testCases[1].expected))
			})
		})
	})
})
