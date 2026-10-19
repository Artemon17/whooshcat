package domain

type Helmet struct {
	InventoryItem
}

func NewHelmet(number, name string) (*Helmet, error) {
	num, err := NewInventoryNumber(number)
	if err != nil {
		return nil, err
	}
	return &Helmet{InventoryItem: NewInventoryItem(num, name)}, nil
}

var _ Inventoriable = (*Helmet)(nil)
