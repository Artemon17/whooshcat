package domain

type ChargingStation struct {
	InventoryItem
	dailyEnergy float64
}

func NewChargingStation(number, name string, energy float64) (*ChargingStation, error) {
	num, err := NewInventoryNumber(number)
	if err != nil {
		return nil, err
	}
	if energy < 0 {
		return nil, ErrNegativeEnergy
	}
	return &ChargingStation{
		InventoryItem: NewInventoryItem(num, name),
		dailyEnergy:   energy,
	}, nil
}

func (c *ChargingStation) DailyEnergyKWh() float64 { return c.dailyEnergy }

var (
	_ Inventoriable  = (*ChargingStation)(nil)
	_ EnergyConsumer = (*ChargingStation)(nil)
)
