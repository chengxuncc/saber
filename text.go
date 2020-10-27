package saber

import (
	"fmt"
	"io"
)

func Print(a ...interface{}) *Compound {
	return Do().Print(a...)
}

func (c *Compound) Print(a ...interface{}) *Compound {
	cmd := c.Do()
	cmd.Call = func(c *Command) error {
		_, err := io.Copy(c.Stdout, c.Stdin)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(c.Stdout, a...)
		return err
	}
	return c
}

func Println(a ...interface{}) *Compound {
	return Do().Println(a...)
}

func (c *Compound) Println(a ...interface{}) *Compound {
	cmd := c.Do()
	cmd.Call = func(c *Command) error {
		_, err := io.Copy(c.Stdout, c.Stdin)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(c.Stdout, a...)
		return err
	}
	return c
}

func Printf(format string, a ...interface{}) *Compound {
	return Do().Printf(format, a...)
}

func (c *Compound) Printf(format string, a ...interface{}) *Compound {
	cmd := c.Do()
	cmd.Call = func(c *Command) error {
		_, err := io.Copy(c.Stdout, c.Stdin)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(c.Stdout, format, a...)
		return err
	}
	return c
}
