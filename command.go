package saber

import (
	"io"
)

type CallFunc func(cmd *Command) error

type Command struct {
	Compound  *Compound
	Stdin     io.ReadCloser
	Stdout    io.WriteCloser
	Stderr    io.WriteCloser
	CallQueue []CallFunc
}

func (c *Command) Queue(fn CallFunc) *Command {
	if fn == nil {
		return c
	}
	c.CallQueue = append(c.CallQueue, fn)
	return c
}

func (c *Command) SetStdin(stream io.ReadCloser) error {
	if c.Stdin != nil {
		err := c.Stdin.Close()
		if err != nil {
			return err
		}
	}
	c.Stdin = stream
	return nil
}

func (c *Command) SetStdout(stream io.WriteCloser) error {
	if c.Stdout != nil {
		err := c.Stdout.Close()
		if err != nil {
			return err
		}
	}
	c.Stdout = stream
	return nil
}

func (c *Command) SetStderr(stream io.WriteCloser) error {
	if c.Stderr != nil {
		err := c.Stderr.Close()
		if err != nil {
			return err
		}
	}
	c.Stderr = stream
	return nil
}
