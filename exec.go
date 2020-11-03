package saber

import (
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func Exec(name string, arg ...string) *Compound {
	return Do().Exec(name, arg...)
}

func (c *Compound) Exec(name string, arg ...string) *Compound {
	return c.Next(func(c *Command) error {
		cmd := exec.Command(name, arg...)
		cmd.Stdin = c.Stdin
		cmd.Stdout = c.Stdout
		cmd.Stderr = c.Stderr
		return cmd.Run()
	})
}

func Eval(cmd string) *Compound {
	return Do().Eval(cmd)
}

func (c *Compound) Eval(cmd string) *Compound {
	return c.Next(func(c *Command) error {
		params, err := shellwords.Parse(cmd)
		if err != nil {
			return err
		}
		if len(params) == 0 {
			return nil
		}
		cmd := exec.Command(params[0], params[1:]...)
		cmd.Stdin = c.Stdin
		cmd.Stdout = c.Stdout
		cmd.Stderr = c.Stderr
		return cmd.Run()
	})
}
