package fleet

import (
	"whooshcat/checkup"
	"whooshcat/domain"
	"whooshcat/rules"
)

type Fleet struct {
	items           []domain.Inventoriable
	transports      []domain.Transport
	energyConsumers []domain.EnergyConsumer
	service         checkup.Inspector
}

func NewFleet(sc checkup.Inspector) *Fleet {
	return &Fleet{service: sc}
}

func (f *Fleet) AcceptTransport(t domain.Transport) (bool, error) {
	ok, err := f.service.Inspect(t)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	f.add(t)
	return true, nil
}

func (f *Fleet) AddItem(i domain.Inventoriable) {
	f.add(i)
}

func (f *Fleet) add(i domain.Inventoriable) {
	f.items = append(f.items, i)
	if t, ok := i.(domain.Transport); ok {
		f.transports = append(f.transports, t)
	}
	if ec, ok := i.(domain.EnergyConsumer); ok {
		f.energyConsumers = append(f.energyConsumers, ec)
	}
}

func (f *Fleet) CountItems() int      { return len(f.items) }
func (f *Fleet) CountTransports() int { return len(f.transports) }

func (f *Fleet) TotalDailyEnergyKWh() float64 {
	var total float64
	for _, ec := range f.energyConsumers {
		total += ec.DailyEnergyKWh()
	}
	return total
}

func (f *Fleet) BeginnerFriendly(p rules.BeginnerPolicy) []domain.Transport {
	var result []domain.Transport
	for _, t := range f.transports {
		if p.IsSuitable(t) {
			result = append(result, t)
		}
	}
	return result
}

func (f *Fleet) AllItems() []domain.Inventoriable {
	return f.items
}
