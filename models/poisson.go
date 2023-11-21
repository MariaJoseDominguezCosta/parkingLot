package models

import "gonum.org/v1/gonum/stat/distuv"

type Poisson struct {
	*distuv.Poisson
}

func NewPoisson() *Poisson {
	return &Poisson{
		Poisson: &distuv.Poisson{
		},
	}
}
func (pd *Poisson) Get(lambda float64) float64 {
	Poisson := distuv.Poisson{
		Lambda: lambda,
		Src: nil,
	}
	return Poisson.Rand()
}