package cli

import (
	"bufio"
	"errors"
	"io"
)

// Command — один пункт меню
type Command interface {
	Key() string
	Title() string
	Run(in *bufio.Reader, out io.Writer) error
}

// ErrExit — специальная ошибка
var ErrExit = errors.New("exit")
