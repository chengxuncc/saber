package saber

import (
	"fmt"
	"os"
)

func Pwd() *Compound {
	return Do().Pwd()
}

func (c *Compound) Pwd() *Compound {
	c.Commands = append(c.Commands,
		func(c *Compound) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(c.Stdout, wd)
			return err
		})
	return c
}
