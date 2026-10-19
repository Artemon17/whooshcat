package rules

import (
	"testing"

	"whooshcat/domain"
)

func TestSimplicityThreshold_RejectsBelow(t *testing.T) {
	p := SimplicityThreshold{Min: 6}
	s, _ := domain.NewScooter("S-1", "X", 1.0, 5)

	if p.IsSuitable(s) {
		t.Fatal("простота 5 не должна подходить")
	}
}

func TestSimplicityThreshold_AcceptsAtThreshold(t *testing.T) {
	p := SimplicityThreshold{Min: 6}
	s, _ := domain.NewScooter("S-1", "X", 1.0, 6)

	if !p.IsSuitable(s) {
		t.Fatal("простота 6 должна подходить")
	}
}

func TestSimplicityThreshold_AcceptsAbove(t *testing.T) {
	p := SimplicityThreshold{Min: 6}
	s, _ := domain.NewScooter("S-1", "X", 1.0, 10)

	if !p.IsSuitable(s) {
		t.Fatal("простота 10 должна подходить")
	}
}
