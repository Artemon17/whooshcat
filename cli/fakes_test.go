package cli

import "whooshcat/domain"

type fakeInspector struct {
	accept bool
	err    error
}

func (f fakeInspector) Inspect(t domain.Transport) (bool, error) {
	return f.accept, f.err
}
