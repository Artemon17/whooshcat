package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (a *App) readString(prompt string) (string, error) {
	fmt.Fprint(a.out, prompt)
	line, err := a.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func (a *App) printError(err error) {
	fmt.Fprintf(a.out, "Ошибка: %v\n", err)
}

func readString(in *bufio.Reader, out io.Writer, prompt string) (string, error) {
	fmt.Fprint(out, prompt)
	line, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func readInt(in *bufio.Reader, out io.Writer, prompt string) (int, error) {
	s, err := readString(in, out, prompt)
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("ожидалось целое число, получено %q", s)
	}
	return n, nil
}

func readFloat(in *bufio.Reader, out io.Writer, prompt string) (float64, error) {
	s, err := readString(in, out, prompt)
	if err != nil {
		return 0, err
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("ожидалось число, получено %q", s)
	}
	return v, nil
}
