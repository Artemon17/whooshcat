package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type App struct {
	commands []Command
	reader   *bufio.Reader
	out      io.Writer
}

func NewApp(commands []Command, in io.Reader, out io.Writer) *App {
	return &App{
		commands: commands,
		reader:   bufio.NewReader(in),
		out:      out,
	}
}

func (a *App) Run() error {
	for {
		a.printMenu()

		key, err := a.readString("> ")
		if err != nil {
			return err
		}

		err = a.dispatch(key)
		if errors.Is(err, ErrExit) {
			fmt.Fprintln(a.out, "Выход.")
			return nil
		}
		if err != nil {
			a.printError(err)
		}
	}
}

func (a *App) printMenu() {
	fmt.Fprintln(a.out, "")
	fmt.Fprintln(a.out, "=== ВышКат: учёт парка ===")
	for _, c := range a.commands {
		fmt.Fprintf(a.out, "%s. %s\n", c.Key(), c.Title())
	}
}

func (a *App) dispatch(key string) error {
	for _, c := range a.commands {
		if c.Key() == key {
			return c.Run(a.reader, a.out)
		}
	}
	fmt.Fprintf(a.out, "Неизвестный пункт: %q\n", key)
	return nil
}
