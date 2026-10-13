package rules

import "whooshcat/domain"

type BeginnerPolicy interface {
	IsSuitable(t domain.Transport) bool
}

type SimplicityThreshold struct {
	Min int
}

func (p SimplicityThreshold) IsSuitable(t domain.Transport) bool {
	return t.BeginnerSimplicity() >= p.Min
}
