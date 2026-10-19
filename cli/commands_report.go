package cli

import (
	"bufio"
	"fmt"
	"io"

	"whooshcat/fleet"
)

type ReportCommand struct {
	fleet *fleet.Fleet
}

func NewReportCommand(f *fleet.Fleet) *ReportCommand {
	return &ReportCommand{fleet: f}
}

func (c *ReportCommand) Key() string { return "3" }
func (c *ReportCommand) Title() string {
	return "Отчёт: количество и суммарная энергия"
}

func (c *ReportCommand) Run(in *bufio.Reader, out io.Writer) error {
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "Всего объектов на балансе: %d\n", c.fleet.CountItems())
	fmt.Fprintf(out, "Из них транспорта: %d\n", c.fleet.CountTransports())
	fmt.Fprintf(out, "Суммарный суточный расход: %.3f кВт·ч\n", c.fleet.TotalDailyEnergyKWh())
	return nil
}
