package checkup

import "whooshcat/domain"

type RealInspector struct {
	minSimplicity int
}

func NewRealInspector(minSimplicity int) *RealInspector {
	return &RealInspector{minSimplicity: minSimplicity}
}

func (r *RealInspector) Inspect(t domain.Transport) (bool, error) {
	if t.BeginnerSimplicity() < r.minSimplicity {
		return false, nil
	}
	return true, nil
}
