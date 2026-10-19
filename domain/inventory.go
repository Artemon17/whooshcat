package domain

type InventoryNumber struct {
	value string
}

func NewInventoryNumber(v string) (InventoryNumber, error) {
	if v == "" {
		return InventoryNumber{}, ErrEmptyInventoryNumber
	}
	return InventoryNumber{value: v}, nil
}

func (n InventoryNumber) String() string { return n.value }

type Inventoriable interface {
	Number() InventoryNumber
	Name() string
}
