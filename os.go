package saber

import (
	"fmt"
	"os"
)

func PWD() *Compound {
	return Do().PWD()
}

func (c *Compound) PWD() *Compound {
	c.Commands = append(c.Commands, &Command{
		Call: func(c *Command) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(c.Stdout, wd)
			return err
		},
	})
	return c
}
