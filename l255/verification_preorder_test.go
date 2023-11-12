package l255_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"kumarkunalbd/datastructure/l255"
)

var _ = Describe("VerificationPreorder", func() {
	Describe("A pre order traversal - SUCCESS", func() {
		Context("When the preorder traversal is given", func() {
			preorder := []int{9, 2, 1, 4, 3, 5, 8, 7, 15, 11, 17}
			preorder1 := []int{9, 2, 1, 4, 3, 5, 8, 7, 11, 15, 17} // same contents as abive but different preorder
			It("It should be a valid preorder and should retrun true", func() {
				ans := l255.VerifyPreorder(preorder)
				Expect(ans).To(BeTrue())
			})

			It("It should be a valid preorder and should retrun true", func() {
				ans := l255.VerifyPreorder(preorder1)
				Expect(ans).Should(BeTrue())
			})
		})
	})

	Describe("A wrong pre order traversal - FAILURE", func() {
		Context("When the preorder traversal is given", func() {
			preorder := []int{9, 2, 1, 4, 3, 5, 8, 7, 15, 17, 11}
			It("It should be not be a valid preorder and should retrun false", func() {
				ans := l255.VerifyPreorder(preorder)
				Expect(ans).ShouldNot(BeTrue())
			})
		})
	})
})
