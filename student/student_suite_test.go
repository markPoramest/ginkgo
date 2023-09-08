package student_test

import (
	"testing"

	"ginkgo/student"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStudent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Student Suite")
}

var _ = Describe("Test Student", func() {
	var std *student.Student

	BeforeEach(func() {
		student.IDRunningNumber = 0
	})

	When("new student", func() {
		It("should got a new student", func() {
			std = student.New("Saharat", "Waenthong", student.Science)

			Expect(std.Id()).To(Equal(int64(1)))
			Expect(std.FirstName()).To(Equal("Saharat"))
			Expect(std.LastName()).To(Equal("Waenthong"))
			Expect(std.Major()).To(Equal(student.Science))
		})

		It("should got a new student with id 1", func() {
			std = student.New("Saharat", "Waenthong", student.Science)

			Expect(std.Id()).To(Equal(int64(1)))
			Expect(std.FirstName()).To(Equal("Saharat"))
			Expect(std.LastName()).To(Equal("Waenthong"))
			Expect(std.Major()).To(Equal(student.Science))
		})
	})

	When("Set GPA", func() {
		It("score 81 should got 4.0", func() {
			std = student.New("Saharat", "Waenthong", student.Science)

			err := std.SetGPA(81)

			Expect(std.GPA()).To(Equal(4.0))
			Expect(err).Error().Should(BeNil())
		})
		It("score 52 should got 1.0", func() {
			std = student.New("Saharat", "Waenthong", student.Science)

			err := std.SetGPA(52)

			Expect(std.GPA()).To(Equal(1.0))
			Expect(err).NotTo(HaveOccurred())
			Expect(err).Error().Should(BeNil())
		})

		It("score 102 should got error", func() {
			std = student.New("Saharat", "Waenthong", student.Science)

			err := std.SetGPA(102)

			Expect(std.GPA()).To(Equal(0.0))
			Expect(err).To(HaveOccurred())
			Expect(err).Error().To(MatchError("score must be between 0 and 100"))
		})
	})
})
