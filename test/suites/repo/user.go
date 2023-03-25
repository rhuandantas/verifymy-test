package repo

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Checking books out of the library", Label("library"), func() {
	BeforeEach(func() {
	})

	When("the library has the book in question", func() {
		BeforeEach(func(ctx SpecContext) {
		})

		Context("and the book is available", func() {
			It("lends it to the reader", func(ctx SpecContext) {
				Expect("").To(Equal("valjean"))
			}, SpecTimeout(time.Second*5))
		})

		Context("but the book has already been checked out", func() {
		})
	})
})
