package quiz

import (
	"errors"
	"fmt"

	"ginkgo/student"
)

type Quiz struct {
	StudentRepo student.Storer
}

func (c *Quiz) VerifyScore(scores []float64, minimumPassScore float64) (string, error) {
	summary := 0.0
	for _, v := range scores {
		if v < minimumPassScore {
			return "", errors.New("not enough score")
		}
		summary += v
	}
	grade, err := c.StudentRepo.CalculateGPA(summary)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("this student has %.2f score and got grade %.1f", summary, grade), nil
}
