package cli

import (
	"bufio"
	"fmt"
	"io"
)

type stubCommand struct {
	key   string
	title string
}

func (c stubCommand) Key() string   { return c.key }
func (c stubCommand) Title() string { return c.title }
func (c stubCommand) Run(in *bufio.Reader, out io.Writer) error {
	fmt.Fprintf(out, "[%s] ещё не реализовано\n", c.title)
	return nil
}

type exitCommand struct{}

func (exitCommand) Key() string   { return "0" }
func (exitCommand) Title() string { return "Выход" }
func (exitCommand) Run(in *bufio.Reader, out io.Writer) error {
	return ErrExit
}

// StubCommands — временный набор команд. Позже заменим на реальные.
func StubCommands() []Command {
	return []Command{
		stubCommand{"1", "Принять технику (с техосмотром)"},
		stubCommand{"2", "Добавить вещь на баланс"},
		stubCommand{"3", "Отчёт: количество и суммарная энергия"},
		stubCommand{"4", "Список техники для новичков"},
		stubCommand{"5", "Показать всё с инвентарными номерами"},
		exitCommand{},
	}
}
