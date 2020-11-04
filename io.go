package saber

import (
	"fmt"
)

func Echo(a ...interface{}) *Compound {
	return Do().Echo(a...)
}

func (c *Compound) Echo(a ...interface{}) *Compound {
	return c.Next(func(c *Command) error {
		_, err := fmt.Fprintln(c.GetStdout(), a...)
		return err
	})
}

func Echon(a ...interface{}) *Compound {
	return Do().Echon(a...)
}

// echo without newline
func (c *Compound) Echon(a ...interface{}) *Compound {
	return c.Next(func(cmd *Command) error {
		_, err := fmt.Fprint(cmd.GetStdout(), a...)
		return err
	})
}

func Printf(format string, a ...interface{}) *Compound {
	return Do().Printf(format, a...)
}

func (c *Compound) Printf(format string, a ...interface{}) *Compound {
	return c.Next(func(cmd *Command) error {
		_, err := fmt.Fprintf(cmd.GetStdout(), format, a...)
		return err
	})
}

func (c *Compound) Combine() *Compound {
	c.Combined = true
	return c
}
