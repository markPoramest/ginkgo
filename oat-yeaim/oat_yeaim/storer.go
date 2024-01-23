package oat_yeaim

//go:generate mockgen -source=storer.go -destination=mock_storer/storer.go

type Storer interface {
	CalculateOatYeaim(yeaim int) string
}
