package cli

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

type echoCommand struct {
	key   string
	title string
}

func (c echoCommand) Key() string   { return c.key }
func (c echoCommand) Title() string { return c.title }
func (c echoCommand) Run(in *bufio.Reader, out io.Writer) error {
	_, _ = io.WriteString(out, "echo:"+c.key+"\n")
	return nil
}

type failCommand struct {
	key   string
	title string
	err   error
}

func (c failCommand) Key() string   { return c.key }
func (c failCommand) Title() string { return c.title }
func (c failCommand) Run(in *bufio.Reader, out io.Writer) error {
	return c.err
}

func TestApp_ExitImmediately(t *testing.T) {
	in := strings.NewReader("0\n")
	out := &bytes.Buffer{}

	app := NewApp([]Command{NewExitCommand()}, in, out)

	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "Выход") {
		t.Fatalf("ожидалось прощание, получили:\n%s", out.String())
	}
}

func TestApp_DispatchesCommand(t *testing.T) {
	// "1" — команда echo, "0" — выход
	in := strings.NewReader("1\n0\n")
	out := &bytes.Buffer{}

	app := NewApp([]Command{
		echoCommand{"1", "Эхо"},
		NewExitCommand(),
	}, in, out)

	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "echo:1") {
		t.Fatalf("ожидался вызов echoCommand, получили:\n%s", out.String())
	}
}

func TestApp_UnknownCommand(t *testing.T) {
	in := strings.NewReader("99\n0\n")
	out := &bytes.Buffer{}

	app := NewApp([]Command{NewExitCommand()}, in, out)

	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "Неизвестный пункт") {
		t.Fatalf("ожидалось сообщение о неизвестном пункте, получили:\n%s", out.String())
	}
}

func TestApp_CommandErrorIsPrinted(t *testing.T) {
	sentinel := errors.New("тестовая ошибка")
	in := strings.NewReader("1\n0\n")
	out := &bytes.Buffer{}

	app := NewApp([]Command{
		failCommand{"1", "Падает", sentinel},
		NewExitCommand(),
	}, in, out)

	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "тестовая ошибка") {
		t.Fatalf("ожидалось сообщение об ошибке, получили:\n%s", out.String())
	}
}

func TestApp_EOFReturnsError(t *testing.T) {
	// пустой ввод — ReadString вернёт io.EOF
	in := strings.NewReader("")
	out := &bytes.Buffer{}

	app := NewApp([]Command{NewExitCommand()}, in, out)

	err := app.Run()
	if err == nil {
		t.Fatal("ожидалась ошибка при EOF")
	}
	if !errors.Is(err, io.EOF) {
		t.Fatalf("ожидался io.EOF, получили %v", err)
	}
}

func TestApp_MenuPrintedFromCommands(t *testing.T) {
	in := strings.NewReader("0\n")
	out := &bytes.Buffer{}

	app := NewApp([]Command{
		echoCommand{"1", "Первая"},
		echoCommand{"2", "Вторая"},
		NewExitCommand(),
	}, in, out)

	if err := app.Run(); err != nil {
		t.Fatal(err)
	}

	got := out.String()
	if !strings.Contains(got, "1. Первая") {
		t.Fatalf("ожидался пункт 1, получили:\n%s", got)
	}
	if !strings.Contains(got, "2. Вторая") {
		t.Fatalf("ожидался пункт 2, получили:\n%s", got)
	}
	if !strings.Contains(got, "0. Выход") {
		t.Fatalf("ожидался пункт 0, получили:\n%s", got)
	}
}
