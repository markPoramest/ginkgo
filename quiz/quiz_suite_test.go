package quiz_test

import (
	"errors"
	"ginkgo/quiz"
	"ginkgo/student/mock_storer"
	"github.com/golang/mock/gomock"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestQuiz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Quiz Suite")
}

var _ = Describe("TestQuiz", func() {
	var ctrl *gomock.Controller

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	When("score is pass minimum", func() {
		It("should successfully", func() {
			scores := []float64{16, 13, 14, 13.5, 17}
			minimumScore := 12.0

			mockStd := mock_student.NewMockStorer(ctrl)
			mockStd.EXPECT().CalculateGPA(73.5).Return(3.0, nil)

			qz := quiz.Quiz{
				StudentRepo: mockStd,
			}

			actualTxt, err := qz.VerifyScore(scores, minimumScore)

			Expect(err).NotTo(HaveOccurred())
			Expect(actualTxt).To(Equal("this student has 73.50 score and got grade 3.0"))
		})
	})

	When("score is not pass minimum", func() {
		It("should error", func() {
			scores := []float64{16, 13, 10, 13.5, 17}
			minimumScore := 12.0

			qz := quiz.Quiz{}

			actualTxt, err := qz.VerifyScore(scores, minimumScore)

			Expect(err).To(MatchError("not enough score"))
			Expect(actualTxt).To(Equal(""))
		})
	})

	When("calculate gpa failed", func() {
		It("should error", func() {
			scores := []float64{16, 13, 14, 13.5, 17}
			minimumScore := 12.0

			mockStd := mock_student.NewMockStorer(ctrl)
			mockStd.EXPECT().CalculateGPA(73.5).Return(0.0, errors.New("gpa failed"))

			qz := quiz.Quiz{
				StudentRepo: mockStd,
			}

			actualTxt, err := qz.VerifyScore(scores, minimumScore)

			Expect(err).To(MatchError("gpa failed"))
			Expect(actualTxt).To(Equal(""))
		})
	})
})
