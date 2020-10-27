package saber

import (
	"io"
)

type CommandCall func(c *Command) error

type Command struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	Call   CommandCall
}
