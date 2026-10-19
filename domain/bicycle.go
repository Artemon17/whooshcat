package domain

type Bicycle struct {
	InventoryItem
	simplicity int
}

func NewBicycle(number, name string, simplicity int) (*Bicycle, error) {
	num, err := NewInventoryNumber(number)
	if err != nil {
		return nil, err
	}
	if simplicity < 1 || simplicity > 10 {
		return nil, ErrInvalidSimplicity
	}
	return &Bicycle{
		InventoryItem: NewInventoryItem(num, name),
		simplicity:    simplicity,
	}, nil
}

func (b *Bicycle) BeginnerSimplicity() int { return b.simplicity }

var _ Transport = (*Bicycle)(nil) // проверка соответствия интерфейсу без создания объекта
