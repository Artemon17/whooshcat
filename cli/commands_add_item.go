package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"whooshcat/domain"
	"whooshcat/fleet"
)

type AddItemCommand struct {
	fleet *fleet.Fleet
}

func NewAddItemCommand(f *fleet.Fleet) *AddItemCommand {
	return &AddItemCommand{fleet: f}
}

func (c *AddItemCommand) Key() string   { return "2" }
func (c *AddItemCommand) Title() string { return "Добавить вещь на баланс" }

func (c *AddItemCommand) Run(in *bufio.Reader, out io.Writer) error {
	kind, err := c.chooseKind(in, out)
	if err != nil {
		return err
	}
	if kind == "" {
		fmt.Fprintln(out, "Отменено.")
		return nil
	}

	common, err := readItemCommon(in, out)
	if err != nil {
		return err
	}

	item, err := c.buildItem(kind, common, in, out)
	if err != nil {
		return err
	}

	c.fleet.AddItem(item)
	fmt.Fprintf(out, "Вещь %s (%s) добавлена на баланс.\n",
		item.Name(), item.Number().String())
	return nil
}

// chooseKind печатает подменю и возвращает выбранный тип.
// Пустая строка — отмена.
func (c *AddItemCommand) chooseKind(in *bufio.Reader, out io.Writer) (string, error) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Выберите тип вещи:")
	fmt.Fprintln(out, "  1. Шлем")
	fmt.Fprintln(out, "  2. Зарядная станция")
	fmt.Fprintln(out, "  3. Док-станция")
	fmt.Fprintln(out, "  0. Отмена")

	choice, err := readString(in, out, "> ")
	if err != nil {
		return "", err
	}

	switch strings.TrimSpace(choice) {
	case "1", "шлем", "helmet":
		return "helmet", nil
	case "2", "зарядная станция", "charging":
		return "charging", nil
	case "3", "док-станция", "dock":
		return "dock", nil
	case "0", "":
		return "", nil
	default:
		fmt.Fprintf(out, "Неизвестный тип: %q\n", choice)
		return "", nil
	}
}

type itemCommon struct {
	number string
	name   string
}

func readItemCommon(in *bufio.Reader, out io.Writer) (itemCommon, error) {
	number, err := readString(in, out, "Инвентарный номер: ")
	if err != nil {
		return itemCommon{}, err
	}
	name, err := readString(in, out, "Название: ")
	if err != nil {
		return itemCommon{}, err
	}
	return itemCommon{number: number, name: name}, nil
}

// buildItem читает специфичные поля и создаёт вещь.
func (c *AddItemCommand) buildItem(
	kind string,
	common itemCommon,
	in *bufio.Reader,
	out io.Writer,
) (domain.Inventoriable, error) {
	switch kind {
	case "helmet":
		return domain.NewHelmet(common.number, common.name)

	case "charging":
		energy, err := readFloat(in, out, "Суточный расход, кВт·ч: ")
		if err != nil {
			return nil, err
		}
		return domain.NewChargingStation(common.number, common.name, energy)

	case "dock":
		return domain.NewDockStation(common.number, common.name)
	}
	return nil, fmt.Errorf("неизвестный тип вещи: %q", kind)
}
