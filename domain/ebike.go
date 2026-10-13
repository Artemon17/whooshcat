package domain

type EBike struct {
	InventoryItem
	dailyEnergy float64
	simplicity  int
}

func NewEBike(number, name string, energy float64, simplicity int) (*EBike, error) {
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
	return &EBike{
		InventoryItem: NewInventoryItem(num, name),
		dailyEnergy:   energy,
		simplicity:    simplicity,
	}, nil
}

func (e *EBike) DailyEnergyKWh() float64 { return e.dailyEnergy }
func (e *EBike) BeginnerSimplicity() int { return e.simplicity }

var _ ElectricTransport = (*EBike)(nil)
