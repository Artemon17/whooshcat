package domain

type InventoryItem struct {
	number InventoryNumber
	name   string
}

func NewInventoryItem(number InventoryNumber, name string) InventoryItem {
	return InventoryItem{number: number, name: name}
}

func (i InventoryItem) Number() InventoryNumber { return i.number }
func (i InventoryItem) Name() string            { return i.name }
