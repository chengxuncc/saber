package saber

import (
	"fmt"
	"io"
	"os"
)

type Std struct {
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
	Stderr io.WriteCloser
}

func (s *Std) SetStdin(stream io.ReadCloser) (err error) {
	if s == nil {
		return
	}
	if s.Stdin != nil && s.Stdin != os.Stdin {
		err = s.Stdin.Close()
		if err != nil {
			return
		}
	}
	s.Stdin = stream
	return
}

func (s *Std) SetStdout(stream io.WriteCloser) (err error) {
	if s == nil {
		return
	}
	if s.Stdout != nil && s.Stdout != os.Stdout {
		err = s.Stdout.Close()
		if err != nil {
			return
		}
	}
	s.Stdout = stream
	return
}

func (s *Std) SetStderr(stream io.WriteCloser) (err error) {
	if s == nil {
		return
	}
	if s.Stderr != nil && s.Stderr != os.Stderr {
		err = s.Stderr.Close()
		if err != nil {
			return
		}
	}
	s.Stderr = stream
	return
}

func Echo(a ...interface{}) *Compound {
	return Do().Echo(a...)
}

func (c *Compound) Echo(a ...interface{}) *Compound {
	return c.Next(func(c *Command) error {
		if c.Stdin != nil {
			_, err := io.Copy(c.Stdout, c.Stdin)
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintln(c.Stdout, a...)
		return err
	})
}

func Echon(a ...interface{}) *Compound {
	return Do().Echon(a...)
}

// echo without newline
func (c *Compound) Echon(a ...interface{}) *Compound {
	return c.Next(func(c *Command) error {
		if c.Stdin != nil {
			_, err := io.Copy(c.Stdout, c.Stdin)
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprint(c.Stdout, a...)
		return err
	})
}

func Printf(format string, a ...interface{}) *Compound {
	return Do().Printf(format, a...)
}

func (c *Compound) Printf(format string, a ...interface{}) *Compound {
	return c.Next(func(cmd *Command) error {
		if cmd.Stdin != nil {
			_, err := io.Copy(cmd.Stdout, cmd.Stdin)
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintf(cmd.Stdout, format, a...)
		return err
	})
}

func (c *Compound) Combine() *Compound {
	return c.Queue(func(cmd *Command) error {
		err := cmd.SetStderr(cmd.Stdout)
		if err != nil {
			return err
		}
		return nil
	})
}
