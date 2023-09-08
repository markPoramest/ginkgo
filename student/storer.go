package student

//go:generate mockgen -source=storer.go -destination=mock_storer/storer.go

type Storer interface {
	CalculateGPA(score float64) (float64, error)
}
