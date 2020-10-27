package saber

import (
	"fmt"
)

func Print(a ...interface{}) *Compound {
	return Do().Print(a...)
}

func (c *Compound) Print(a ...interface{}) *Compound {
	c.Commands = append(c.Commands,
		func(c *Compound) error {
			_, err := fmt.Fprint(c.Stdout, a...)
			return err
		})
	return c
}

func Println(a ...interface{}) *Compound {
	return Do().Print(a...)
}

func (c *Compound) Println(a ...interface{}) *Compound {
	c.Commands = append(c.Commands,
		func(c *Compound) error {
			_, err := fmt.Fprintln(c.Stdout, a...)
			return err
		})
	return c
}

func Printf(format string, a ...interface{}) *Compound {
	return Do().Printf(format, a...)
}

func (c *Compound) Printf(format string, a ...interface{}) *Compound {
	c.Commands = append(c.Commands,
		func(c *Compound) error {
			_, err := fmt.Fprintf(c.Stdout, format, a...)
			return err
		})
	return c
}
