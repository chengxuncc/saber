package saber

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Pwd() *Compound {
	return Do().Pwd()
}

func (c *Compound) Pwd() *Compound {
	cmd := c.Do()
	cmd.Call = func(c *Command) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(c.Stdout, wd)
		return err
	}
	return c
}

func Cat(file string) *Compound {
	return Do().Cat(file)
}

func (c *Compound) Cat(file string) *Compound {
	cmd := c.Do()
	cmd.Call = func(c *Command) error {
		if c.Stdin != nil {
			return errors.New("saber: Stdin is already set")
		}
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(c.Stdout, f)
		if err != nil {
			return err
		}
		return nil
	}
	return c
}
