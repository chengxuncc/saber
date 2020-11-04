package saber

import (
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func Exec(name string, arg ...string) *Compound {
	return Do().Exec(name, arg...)
}

func (c *Compound) Exec(name string, arg ...string) *Compound {
	return c.Next(func(cmd *Command) error {
		proc := exec.Command(name, arg...)
		proc.Stdin = cmd.GetStdin()
		proc.Stdout = cmd.GetStdout()
		proc.Stderr = cmd.GetStderr()
		return proc.Run()
	})
}

func Eval(cmd string) *Compound {
	return Do().Eval(cmd)
}

func (c *Compound) Eval(shell string) *Compound {
	return c.Next(func(cmd *Command) error {
		params, err := shellwords.Parse(shell)
		if err != nil {
			return err
		}
		if len(params) == 0 {
			return nil
		}
		proc := exec.Command(params[0], params[1:]...)
		proc.Stdin = cmd.GetStdin()
		proc.Stdout = cmd.GetStdout()
		proc.Stderr = cmd.GetStderr()
		return proc.Run()
	})
}
