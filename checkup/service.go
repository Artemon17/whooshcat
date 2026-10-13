package checkup

import "whooshcat/domain"

type Inspector interface {
	Inspect(t domain.Transport) (bool, error)
}
