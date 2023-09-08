package student

import "errors"

type Major int

var IDRunningNumber int64 = 0

const (
	Science  Major = 0
	Art      Major = 1
	Engineer Major = 2
	Law      Major = 3
)

type Student struct {
	id        int64
	firstName string
	lastName  string
	gpa       float64
	year      int64
	major     Major
}

func (s *Student) Id() int64 {
	return s.id
}

func (s *Student) FirstName() string {
	return s.firstName
}

func (s *Student) LastName() string {
	return s.lastName
}

func (s *Student) GPA() float64 {
	return s.gpa
}

func (s *Student) Year() int64 {
	return s.year
}

func (s *Student) Major() Major {
	return s.major
}

func New(firstName string, lastName string, major Major) *Student {
	return &Student{
		id:        runningStudentIDNumber(),
		firstName: firstName,
		lastName:  lastName,
		major:     major,
		year:      1,
	}
}

func runningStudentIDNumber() int64 {
	IDRunningNumber += 1
	return IDRunningNumber
}

func (s *Student) SetFirstName(name string) {
	s.firstName = name
}

func (s *Student) SetLastName(name string) {
	s.lastName = name
}

func (s *Student) SetMajor(major Major) {
	s.major = major
}

func (s *Student) SetYear(year int64) {
	s.year = year
}

func (s *Student) SetGPA(score float64) error {
	grade, err := CalculateGPA(score)
	if err != nil {
		return err
	}
	s.gpa = grade
	return nil
}

func CalculateGPA(score float64) (float64, error) {
	if score > 100 || score < 0 {
		return 0, errors.New("score must be between 0 and 100")
	}
	if score >= 80 {
		return 4.0, nil
	}
	if score >= 75 {
		return 3.5, nil
	}
	if score >= 70 {
		return 3, nil
	}
	if score >= 65 {
		return 2.5, nil
	}
	if score >= 60 {
		return 2, nil
	}
	if score >= 55 {
		return 1.5, nil
	}
	if score >= 50 {
		return 1, nil
	}
	return 0, nil
}
