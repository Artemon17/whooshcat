package cli

import (
	"bufio"
	"fmt"
	"io"

	"whooshcat/fleet"
)

type AllItemsCommand struct {
	fleet *fleet.Fleet
}

func NewAllItemsCommand(f *fleet.Fleet) *AllItemsCommand {
	return &AllItemsCommand{fleet: f}
}

func (c *AllItemsCommand) Key() string { return "5" }
func (c *AllItemsCommand) Title() string {
	return "Показать всё с инвентарными номерами"
}

func (c *AllItemsCommand) Run(in *bufio.Reader, out io.Writer) error {
	items := c.fleet.AllItems()
	fmt.Fprintln(out, "")
	if len(items) == 0 {
		fmt.Fprintln(out, "На балансе пока ничего нет.")
		return nil
	}
	fmt.Fprintln(out, "Всё на балансе:")
	for _, i := range items {
		fmt.Fprintf(out, "  %s — %s\n", i.Number().String(), i.Name())
	}
	return nil
}
