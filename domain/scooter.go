package domain

type Scooter struct {
	InventoryItem
	dailyEnergy float64
	simplicity  int
}

func NewScooter(number, name string, energy float64, simplicity int) (*Scooter, error) {
	num, err := NewInventoryNumber(number)
	if err != nil {
		return nil, err
	}
	if simplicity < 1 || simplicity > 10 {
		return nil, ErrInvalidSimplicity
	}
	if energy < 0 {
		return nil, ErrNegativeEnergy
	}
	return &Scooter{
		InventoryItem: NewInventoryItem(num, name),
		dailyEnergy:   energy,
		simplicity:    simplicity,
	}, nil
}

func (s *Scooter) DailyEnergyKWh() float64 { return s.dailyEnergy }
func (s *Scooter) BeginnerSimplicity() int { return s.simplicity }

var _ ElectricTransport = (*Scooter)(nil)
