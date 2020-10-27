package saber

import (
	"fmt"
	"io"
)

func Print(a ...interface{}) *Compound {
	return Do().Print(a...)
}

func (c *Compound) Print(a ...interface{}) *Compound {
	c.Commands = append(c.Commands, &Command{
		Call: func(c *Command) error {
			if c.Stdin != nil {
				_, err := io.Copy(c.Stdout, c.Stdin)
				if err != nil {
					return err
				}
			}
			_, err := fmt.Fprint(c.Stdout, a...)
			return err
		},
	})
	return c
}
