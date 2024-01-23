package oat_yeaim_test

import (
	"testing"

	"ginkgo/oat-yeaim/oat_yeaim"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOatYeaim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OatYeaim Suite")
}

var _ = Describe("Oat Yeaim", func() {
	When("input is 0", func() {
		It("should got oat", func() {
			Expect(oat_yeaim.CalculateOatYeaim(0)).To(Equal("oat"))
		})
	})

	When("input is 1", func() {
		It("should got yeaim", func() {
			Expect(oat_yeaim.CalculateOatYeaim(1)).To(Equal("yeaim"))
		})
	})
})
