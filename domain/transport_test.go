package domain

import (
	"errors"
	"testing"
)

func TestNewScooter(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		model   string
		energy  float64
		simpl   int
		wantErr error
	}{
		{"valid", "S-1", "Xiaomi", 1.2, 7, nil},
		{"empty number", "", "Xiaomi", 1.2, 7, ErrEmptyInventoryNumber},
		{"negative energy", "S-1", "Xiaomi", -0.1, 7, ErrNegativeEnergy},
		{"simplicity zero", "S-1", "Xiaomi", 1.2, 0, ErrInvalidSimplicity},
		{"simplicity eleven", "S-1", "Xiaomi", 1.2, 11, ErrInvalidSimplicity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewScooter(tt.number, tt.model, tt.energy, tt.simpl)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ожидалась %v, получили %v", tt.wantErr, err)
			}
		})
	}
}

func TestNewBicycle(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		model   string
		simpl   int
		wantErr error
	}{
		{"valid", "B-1", "Stels", 9, nil},
		{"empty number", "", "Stels", 9, ErrEmptyInventoryNumber},
		{"simplicity zero", "B-1", "Stels", 0, ErrInvalidSimplicity},
		{"simplicity eleven", "B-1", "Stels", 11, ErrInvalidSimplicity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBicycle(tt.number, tt.model, tt.simpl)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ожидалась %v, получили %v", tt.wantErr, err)
			}
		})
	}
}

func TestNewEBike(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		model   string
		energy  float64
		simpl   int
		wantErr error
	}{
		{"valid", "E-1", "Aventon", 2.5, 8, nil},
		{"empty number", "", "Aventon", 2.5, 8, ErrEmptyInventoryNumber},
		{"negative energy", "E-1", "Aventon", -1, 8, ErrNegativeEnergy},
		{"simplicity zero", "E-1", "Aventon", 2.5, 0, ErrInvalidSimplicity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEBike(tt.number, tt.model, tt.energy, tt.simpl)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ожидалась %v, получили %v", tt.wantErr, err)
			}
		})
	}
}

func TestScooter_Methods(t *testing.T) {
	s, _ := NewScooter("S-1", "Xiaomi", 1.2, 7)
	if s.Number().String() != "S-1" {
		t.Fatalf("ожидалось S-1, получили %s", s.Number().String())
	}
	if s.Name() != "Xiaomi" {
		t.Fatalf("ожидалось Xiaomi, получили %s", s.Name())
	}
	if s.DailyEnergyKWh() != 1.2 {
		t.Fatalf("ожидалось 1.2, получили %v", s.DailyEnergyKWh())
	}
	if s.BeginnerSimplicity() != 7 {
		t.Fatalf("ожидалось 7, получили %d", s.BeginnerSimplicity())
	}
}
