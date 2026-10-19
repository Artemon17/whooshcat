package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"whooshcat/domain"
	"whooshcat/fleet"
)

type AcceptTransportCommand struct {
	fleet *fleet.Fleet
}

func NewAcceptTransportCommand(f *fleet.Fleet) *AcceptTransportCommand {
	return &AcceptTransportCommand{fleet: f}
}

func (c *AcceptTransportCommand) Key() string { return "1" }
func (c *AcceptTransportCommand) Title() string {
	return "Принять технику (с техосмотром)"
}

func (c *AcceptTransportCommand) Run(in *bufio.Reader, out io.Writer) error {
	kind, err := c.chooseKind(in, out)
	if err != nil {
		return err
	}
	if kind == "" {
		fmt.Fprintln(out, "Отменено.")
		return nil
	}

	common, err := readTransportCommon(in, out)
	if err != nil {
		return err
	}

	t, err := c.buildTransport(kind, common, in, out)
	if err != nil {
		return err
	}

	ok, err := c.fleet.AcceptTransport(t)
	if err != nil {
		return err
	}
	if ok {
		fmt.Fprintf(out, "Техника %s (%s) принята в парк.\n", t.Name(), t.Number().String())
	} else {
		fmt.Fprintf(out, "Техника %s (%s) отклонена на техосмотре.\n", t.Name(), t.Number().String())
	}
	return nil
}

func (c *AcceptTransportCommand) chooseKind(in *bufio.Reader, out io.Writer) (string, error) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Выберите тип транспорта:")
	fmt.Fprintln(out, "  1. Самокат")
	fmt.Fprintln(out, "  2. Велосипед")
	fmt.Fprintln(out, "  3. Электровелосипед")
	fmt.Fprintln(out, "  0. Отмена")

	choice, err := readString(in, out, "> ")
	if err != nil {
		return "", err
	}

	switch strings.TrimSpace(choice) {
	case "1", "самокат", "scooter":
		return "scooter", nil
	case "2", "велосипед", "bicycle":
		return "bicycle", nil
	case "3", "электровелосипед", "ebike":
		return "ebike", nil
	case "0", "":
		return "", nil
	default:
		fmt.Fprintf(out, "Неизвестный тип: %q\n", choice)
		return "", nil // не ошибка, просто вернёмся в главное меню
	}
}

type transportCommon struct {
	number     string
	name       string
	simplicity int
}

func readTransportCommon(in *bufio.Reader, out io.Writer) (transportCommon, error) {
	number, err := readString(in, out, "Инвентарный номер: ")
	if err != nil {
		return transportCommon{}, err
	}
	name, err := readString(in, out, "Название: ")
	if err != nil {
		return transportCommon{}, err
	}
	simplicity, err := readInt(in, out, "Простота для новичка (1–10): ")
	if err != nil {
		return transportCommon{}, err
	}
	return transportCommon{number: number, name: name, simplicity: simplicity}, nil
}

// buildTransport читает специфичные поля и создаёт транспорт.
func (c *AcceptTransportCommand) buildTransport(
	kind string,
	common transportCommon,
	in *bufio.Reader,
	out io.Writer,
) (domain.Transport, error) {
	switch kind {
	case "scooter":
		energy, err := readFloat(in, out, "Суточный расход, кВт·ч: ")
		if err != nil {
			return nil, err
		}
		return domain.NewScooter(common.number, common.name, energy, common.simplicity)

	case "bicycle":
		return domain.NewBicycle(common.number, common.name, common.simplicity)

	case "ebike":
		energy, err := readFloat(in, out, "Суточный расход, кВт·ч: ")
		if err != nil {
			return nil, err
		}
		return domain.NewEBike(common.number, common.name, energy, common.simplicity)
	}
	return nil, fmt.Errorf("неизвестный тип транспорта: %q", kind)
}
