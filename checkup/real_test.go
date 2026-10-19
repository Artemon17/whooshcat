package checkup

import (
	"testing"

	"whooshcat/domain"
)

func TestRealInspector_AcceptsAtThreshold(t *testing.T) {
	insp := NewRealInspector(3)
	s, _ := domain.NewScooter("S-1", "X", 1.0, 3)

	ok, err := insp.Inspect(s)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("ожидалось принятие при простоте == порогу")
	}
}

func TestRealInspector_RejectsBelow(t *testing.T) {
	insp := NewRealInspector(3)
	s, _ := domain.NewScooter("S-1", "X", 1.0, 2)

	ok, err := insp.Inspect(s)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("ожидалось отклонение при простоте < порога")
	}
}

func TestRealInspector_AcceptsAbove(t *testing.T) {
	insp := NewRealInspector(3)
	s, _ := domain.NewScooter("S-1", "X", 1.0, 7)

	ok, _ := insp.Inspect(s)
	if !ok {
		t.Fatal("ожидалось принятие при простоте > порога")
	}
}
