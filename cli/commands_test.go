package cli

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"whooshcat/domain"
	"whooshcat/fleet"
	"whooshcat/rules"
)

func runCommand(cmd Command, input string) (string, *fleet.Fleet, error) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	_ = f
	in := strings.NewReader(input)
	out := &bytes.Buffer{}
	err := cmd.Run(bufio.NewReader(in), out)
	return out.String(), nil, err
}

func runWithFleet(cmd Command, f *fleet.Fleet, input string) (string, error) {
	in := strings.NewReader(input)
	out := &bytes.Buffer{}
	err := cmd.Run(bufio.NewReader(in), out)
	return out.String(), err
}

func TestAcceptTransportCommand_Scooter(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAcceptTransportCommand(f)

	// подменю: 1 (самокат), номер, имя, простота, энергия
	out, err := runWithFleet(cmd, f, "1\nS-1\nXiaomi\n7\n1.2\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountTransports() != 1 {
		t.Fatalf("ожидался 1 транспорт, получили %d", f.CountTransports())
	}
	if !strings.Contains(out, "принята в парк") {
		t.Fatalf("ожидалось подтверждение, получили:\n%s", out)
	}
}

func TestAcceptTransportCommand_Bicycle(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAcceptTransportCommand(f)

	out, err := runWithFleet(cmd, f, "2\nB-1\nStels\n9\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountTransports() != 1 {
		t.Fatalf("ожидался 1 транспорт, получили %d", f.CountTransports())
	}
	if !strings.Contains(out, "принята в парк") {
		t.Fatalf("ожидалось подтверждение, получили:\n%s", out)
	}
}

func TestAcceptTransportCommand_EBike(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAcceptTransportCommand(f)

	_, err := runWithFleet(cmd, f, "3\nE-1\nAventon\n8\n2.5\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountTransports() != 1 {
		t.Fatalf("ожидался 1 транспорт, получили %d", f.CountTransports())
	}
}

func TestAcceptTransportCommand_Rejected(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: false})
	cmd := NewAcceptTransportCommand(f)

	out, err := runWithFleet(cmd, f, "1\nS-1\nXiaomi\n7\n1.2\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountTransports() != 0 {
		t.Fatal("транспорт не должен попасть в парк")
	}
	if !strings.Contains(out, "отклонена") {
		t.Fatalf("ожидалось отклонение, получили:\n%s", out)
	}
}

func TestAcceptTransportCommand_Cancel(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAcceptTransportCommand(f)

	out, err := runWithFleet(cmd, f, "0\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountTransports() != 0 {
		t.Fatal("после отмены ничего не должно добавиться")
	}
	if !strings.Contains(out, "Отменено") {
		t.Fatalf("ожидалось сообщение об отмене, получили:\n%s", out)
	}
}

func TestAddItemCommand_Helmet(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAddItemCommand(f)

	out, err := runWithFleet(cmd, f, "1\nH-1\nШлем\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountItems() != 1 {
		t.Fatalf("ожидался 1 объект, получили %d", f.CountItems())
	}
	if !strings.Contains(out, "добавлена на баланс") {
		t.Fatalf("ожидалось подтверждение, получили:\n%s", out)
	}
}

func TestAddItemCommand_ChargingStation(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAddItemCommand(f)

	_, err := runWithFleet(cmd, f, "2\nC-1\nChargePoint\n3.5\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.TotalDailyEnergyKWh() != 3.5 {
		t.Fatalf("ожидалось 3.5 кВт·ч, получили %v", f.TotalDailyEnergyKWh())
	}
}

func TestAddItemCommand_DockStation(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAddItemCommand(f)

	_, err := runWithFleet(cmd, f, "3\nD-1\nДок\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountItems() != 1 {
		t.Fatalf("ожидался 1 объект, получили %d", f.CountItems())
	}
}

func TestAddItemCommand_Cancel(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAddItemCommand(f)

	out, err := runWithFleet(cmd, f, "0\n")
	if err != nil {
		t.Fatal(err)
	}
	if f.CountItems() != 0 {
		t.Fatal("после отмены ничего не должно добавиться")
	}
	if !strings.Contains(out, "Отменено") {
		t.Fatalf("ожидалось сообщение об отмене, получили:\n%s", out)
	}
}

func TestReportCommand(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	s, _ := domain.NewScooter("S-1", "Xiaomi", 1.2, 7)
	f.AcceptTransport(s)
	cs, _ := domain.NewChargingStation("C-1", "ChargePoint", 3.5)
	f.AddItem(cs)

	cmd := NewReportCommand(f)
	out := &bytes.Buffer{}
	if err := cmd.Run(bufio.NewReader(strings.NewReader("")), out); err != nil {
		t.Fatal(err)
	}

	got := out.String()
	if !strings.Contains(got, "Всего объектов на балансе: 2") {
		t.Fatalf("ожидалось 2 объекта, получили:\n%s", got)
	}
	if !strings.Contains(got, "Из них транспорта: 1") {
		t.Fatalf("ожидался 1 транспорт, получили:\n%s", got)
	}
	if !strings.Contains(got, "4.700") {
		t.Fatalf("ожидалось 4.700 кВт·ч, получили:\n%s", got)
	}
}

func TestBeginnerListCommand_NotEmpty(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	good, _ := domain.NewScooter("S-1", "Good", 1.0, 7)
	bad, _ := domain.NewScooter("S-2", "Bad", 1.0, 4)
	f.AcceptTransport(good)
	f.AcceptTransport(bad)

	cmd := NewBeginnerListCommand(f, rules.SimplicityThreshold{Min: 6})
	out := &bytes.Buffer{}
	if err := cmd.Run(bufio.NewReader(strings.NewReader("")), out); err != nil {
		t.Fatal(err)
	}

	got := out.String()
	if !strings.Contains(got, "Good") {
		t.Fatalf("ожидался Good, получили:\n%s", got)
	}
	if strings.Contains(got, "Bad") {
		t.Fatalf("Bad не должен быть в списке, получили:\n%s", got)
	}
}

func TestBeginnerListCommand_Empty(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	bad, _ := domain.NewScooter("S-2", "Bad", 1.0, 4)
	f.AcceptTransport(bad)

	cmd := NewBeginnerListCommand(f, rules.SimplicityThreshold{Min: 6})
	out := &bytes.Buffer{}
	cmd.Run(bufio.NewReader(strings.NewReader("")), out)

	if !strings.Contains(out.String(), "Нет техники") {
		t.Fatalf("ожидалось сообщение об отсутствии, получили:\n%s", out.String())
	}
}

func TestAllItemsCommand_NotEmpty(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	s, _ := domain.NewScooter("S-1", "Xiaomi", 1.2, 7)
	f.AcceptTransport(s)
	h, _ := domain.NewHelmet("H-1", "Шлем")
	f.AddItem(h)

	cmd := NewAllItemsCommand(f)
	out := &bytes.Buffer{}
	cmd.Run(bufio.NewReader(strings.NewReader("")), out)

	got := out.String()
	if !strings.Contains(got, "S-1") || !strings.Contains(got, "H-1") {
		t.Fatalf("ожидались оба номера, получили:\n%s", got)
	}
}

func TestAllItemsCommand_Empty(t *testing.T) {
	f := fleet.NewFleet(fakeInspector{accept: true})
	cmd := NewAllItemsCommand(f)

	out := &bytes.Buffer{}
	cmd.Run(bufio.NewReader(strings.NewReader("")), out)

	if !strings.Contains(out.String(), "ничего нет") {
		t.Fatalf("ожидалось сообщение о пустом парке, получили:\n%s", out.String())
	}
}

func TestExitCommand(t *testing.T) {
	cmd := NewExitCommand()
	in := bufio.NewReader(strings.NewReader(""))
	out := &bytes.Buffer{}

	err := cmd.Run(in, out)
	if err != ErrExit {
		t.Fatalf("ожидалась ErrExit, получили %v", err)
	}
}
