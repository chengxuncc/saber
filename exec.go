package saber

import (
	"os/exec"
)

func Cmd(name string, arg ...string) *Compound {
	return Do().Cmd(name, arg...)
}

func (c *Compound) Cmd(name string, arg ...string) *Compound {
	return c.Call(func(c *Command) error {
		cmd := exec.Command(name, arg...)
		cmd.Stdin = c.Stdin
		cmd.Stdout = c.Stdout
		cmd.Stderr = c.Stderr
		return cmd.Run()
	})
}
