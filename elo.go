package elo

import (
	"errors"
	"math"
)

// Result is a Win | Draw | Loss
type Result float64

// Result codes
const (
	Win   = Result(1)
	Draw  = Result(0.5)
	Loose = Result(0)
)

// ELO is the Elo ranking item
type ELO struct {
	RA  float64
	RB  float64
	K   float64
	EA  float64
	EB  float64
	SA  Result
	SB  Result
	RAN float64
	RBN float64
}

// New returns a new ELO
func New(rA, rB, k float64, sa, sb Result) (*ELO, error) {
	if float64(sa)+float64(sb) != 1 {
		return nil, errors.New("invalid result")
	}
	var e = &ELO{
		RA: rA,
		RB: rB,
		K:  k,
		SA: sa,
		SB: sb,
	}
	e.calculate()
	return e, nil
}

func (e *ELO) calculate() {
	qA := math.Pow(10, e.RA/400)
	qB := math.Pow(10, e.RB/400)

	e.EA = qA / (qA + qB)
	e.EB = qB / (qA + qB)

	e.RAN = e.RA + e.K*(float64(e.SA)-e.EA)
	e.RBN = e.RB + e.K*(float64(e.SB)-e.EB)
}
