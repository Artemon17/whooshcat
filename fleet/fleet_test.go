package fleet

import (
	"testing"

	"whooshcat/domain"
	"whooshcat/rules"
)

type fakeInspector struct {
	accept bool
	err    error
}

func (f fakeInspector) Inspect(t domain.Transport) (bool, error) {
	return f.accept, f.err
}

func TestFleet_AcceptsWhenInspectionPasses(t *testing.T) {
	f := NewFleet(fakeInspector{accept: true})
	s, _ := domain.NewScooter("S-1", "Xiaomi", 1.2, 7)

	ok, err := f.AcceptTransport(s)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("ожидалось принятие")
	}
	if f.CountTransports() != 1 {
		t.Fatalf("ожидался 1 транспорт, получили %d", f.CountTransports())
	}
}

func TestFleet_RejectsWhenInspectionFails(t *testing.T) {
	f := NewFleet(fakeInspector{accept: false})
	s, _ := domain.NewScooter("S-1", "Xiaomi", 1.2, 7)

	ok, _ := f.AcceptTransport(s)
	if ok {
		t.Fatal("ожидалось отклонение")
	}
	if f.CountTransports() != 0 {
		t.Fatal("транспорт не должен попасть в парк")
	}
}

func TestFleet_TotalEnergyIncludesScootersAndStations(t *testing.T) {
	f := NewFleet(fakeInspector{accept: true})
	s, _ := domain.NewScooter("S-1", "Xiaomi", 1.2, 7)
	cs, _ := domain.NewChargingStation("C-1", "ChargePoint", 3.5)
	h, _ := domain.NewHelmet("H-1", "Шлем")

	f.AcceptTransport(s)
	f.AddItem(cs)
	f.AddItem(h)

	got := f.TotalDailyEnergyKWh()
	want := 1.2 + 3.5
	if got != want {
		t.Fatalf("ожидалось %.3f, получили %.3f", want, got)
	}
}

func TestFleet_BeginnerFriendlyFiltersByPolicy(t *testing.T) {
	f := NewFleet(fakeInspector{accept: true})
	good, _ := domain.NewScooter("S-1", "Good", 1.0, 7)
	bad, _ := domain.NewScooter("S-2", "Bad", 1.0, 4)
	bike, _ := domain.NewBicycle("B-1", "Stels", 9)

	f.AcceptTransport(good)
	f.AcceptTransport(bad)
	f.AcceptTransport(bike)

	policy := rules.SimplicityThreshold{Min: 6}
	result := f.BeginnerFriendly(policy)

	if len(result) != 2 {
		t.Fatalf("ожидалось 2 подходящих, получили %d", len(result))
	}
}
