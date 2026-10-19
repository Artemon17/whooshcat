package domain

import (
	"errors"
	"testing"
)

func TestNewInventoryNumber_Valid(t *testing.T) {
	n, err := NewInventoryNumber("S-1")
	if err != nil {
		t.Fatal(err)
	}
	if n.String() != "S-1" {
		t.Fatalf("ожидалось S-1, получили %s", n.String())
	}
}

func TestNewInventoryNumber_Empty(t *testing.T) {
	_, err := NewInventoryNumber("")
	if !errors.Is(err, ErrEmptyInventoryNumber) {
		t.Fatalf("ожидалась ErrEmptyInventoryNumber, получили %v", err)
	}
}
