package saber

import (
	"fmt"
	"io"
)

func Echo(a ...interface{}) *Compound {
	return Do().Echo(a...)
}

func (c *Compound) Echo(a ...interface{}) *Compound {
	return c.Call(func(c *Command) error {
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

func EchoN(a ...interface{}) *Compound {
	return Do().EchoN(a...)
}

// echo without newline
func (c *Compound) EchoN(a ...interface{}) *Compound {
	return c.Call(func(c *Command) error {
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
	return c.Call(func(c *Command) error {
		if c.Stdin != nil {
			_, err := io.Copy(c.Stdout, c.Stdin)
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintf(c.Stdout, format, a...)
		return err
	})
}
