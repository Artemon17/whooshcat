package domain

import (
	"errors"
	"testing"
)

func TestNewHelmet(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		model   string
		wantErr error
	}{
		{"valid", "H-1", "Шлем", nil},
		{"empty number", "", "Шлем", ErrEmptyInventoryNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewHelmet(tt.number, tt.model)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ожидалась %v, получили %v", tt.wantErr, err)
			}
		})
	}
}

func TestNewDockStation(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		model   string
		wantErr error
	}{
		{"valid", "D-1", "Док", nil},
		{"empty number", "", "Док", ErrEmptyInventoryNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDockStation(tt.number, tt.model)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ожидалась %v, получили %v", tt.wantErr, err)
			}
		})
	}
}

func TestNewChargingStation(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		model   string
		energy  float64
		wantErr error
	}{
		{"valid", "C-1", "ChargePoint", 3.5, nil},
		{"empty number", "", "ChargePoint", 3.5, ErrEmptyInventoryNumber},
		{"negative energy", "C-1", "ChargePoint", -1, ErrNegativeEnergy},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewChargingStation(tt.number, tt.model, tt.energy)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ожидалась %v, получили %v", tt.wantErr, err)
			}
		})
	}
}

func TestChargingStation_DailyEnergy(t *testing.T) {
	c, _ := NewChargingStation("C-1", "ChargePoint", 3.5)
	if c.DailyEnergyKWh() != 3.5 {
		t.Fatalf("ожидалось 3.5, получили %v", c.DailyEnergyKWh())
	}
}
