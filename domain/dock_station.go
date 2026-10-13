package domain

type DockStation struct {
	InventoryItem
}

func NewDockStation(number, name string) (*DockStation, error) {
	num, err := NewInventoryNumber(number)
	if err != nil {
		return nil, err
	}
	return &DockStation{InventoryItem: NewInventoryItem(num, name)}, nil
}

var _ Inventoriable = (*DockStation)(nil)
