package domain

import "errors"

var (
	ErrEmptyInventoryNumber = errors.New("инвентарный номер не может быть пустым")
	ErrInvalidSimplicity    = errors.New("простота должна быть в диапазоне 1..10")
	ErrNegativeEnergy       = errors.New("расход энергии не может быть отрицательным")
)
