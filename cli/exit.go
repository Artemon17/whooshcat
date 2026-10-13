package cli

import (
	"bufio"
	"io"
)

type ExitCommand struct{}

func NewExitCommand() *ExitCommand { return &ExitCommand{} }

func (c *ExitCommand) Key() string   { return "0" }
func (c *ExitCommand) Title() string { return "Выход" }

func (c *ExitCommand) Run(in *bufio.Reader, out io.Writer) error {
	return ErrExit
}
