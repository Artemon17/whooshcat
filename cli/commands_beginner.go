package cli

import (
	"bufio"
	"fmt"
	"io"

	"whooshcat/fleet"
	"whooshcat/rules"
)

type BeginnerListCommand struct {
	fleet  *fleet.Fleet
	policy rules.BeginnerPolicy
}

func NewBeginnerListCommand(f *fleet.Fleet, p rules.BeginnerPolicy) *BeginnerListCommand {
	return &BeginnerListCommand{fleet: f, policy: p}
}

func (c *BeginnerListCommand) Key() string { return "4" }
func (c *BeginnerListCommand) Title() string {
	return "Список техники для новичков"
}

func (c *BeginnerListCommand) Run(in *bufio.Reader, out io.Writer) error {
	suitable := c.fleet.BeginnerFriendly(c.policy)
	fmt.Fprintln(out, "")
	if len(suitable) == 0 {
		fmt.Fprintln(out, "Нет техники, подходящей новичкам.")
		return nil
	}
	fmt.Fprintln(out, "Техника для новичков:")
	for _, t := range suitable {
		fmt.Fprintf(out, "  %s — %s (простота %d)\n",
			t.Number().String(), t.Name(), t.BeginnerSimplicity())
	}
	return nil
}
