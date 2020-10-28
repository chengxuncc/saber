package saber

import (
	"fmt"
	"io"
	"os"

	"github.com/chengxuncc/saber/internal/x"
)

type Command struct {
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
	Stderr io.WriteCloser
	Call   func(c *Command) error
}

type Compound struct {
	Script   *Script
	Commands []*Command
}

func Do() *Compound {
	return &Compound{
		Commands: make([]*Command, 0, 3),
	}
}

func (c *Compound) Do() *Command {
	cmd := &Command{}
	if len(c.Commands) > 0 {
		lastCmd := c.Commands[len(c.Commands)-1]
		if lastCmd.Stdout == nil {
			cmd.Stdin, lastCmd.Stdout = io.Pipe()
		}
	}
	c.Commands = append(c.Commands, cmd)
	return cmd
}

func (c *Compound) Call(callable func(c *Command) error) *Compound {
	cmd := c.Do()
	cmd.Call = callable
	return c
}

func (c *Compound) Run() {
	err := c.ErrorRun()
	if err != nil {
		x.Must(fmt.Fprintln(os.Stderr, err))
		if c.Script.ExitOnError {
			os.Exit(1)
		}
	}
}

func (c *Compound) ErrorRun() error {
	if c.Script == nil {
		c.Script = Main
	}
	count := len(c.Commands)
	errs := make(chan error, count)
	for i := 0; i < count; i++ {
		cmd := c.Commands[i]
		if cmd.Stdout == nil {
			cmd.Stdout = os.Stdout
		}
		go func() {
			errs <- cmd.Call(cmd)
			if cmd.Stdout != os.Stdout {
				_ = cmd.Stdout.Close()
			}
		}()
	}
	for i := 0; i < count; i++ {
		err := <-errs
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compound) Output() string {
	res, err := c.ErrorOutput()
	if err != nil {
		x.Must(fmt.Fprintln(os.Stderr, err))
		if c.Script.ExitOnError {
			os.Exit(1)
		}
	}
	return res
}

func (c *Compound) ErrorOutput() (string, error) {
	out := &x.Buffer{}
	cmdCount := len(c.Commands)
	if cmdCount > 0 {
		c.Commands[cmdCount-1].Stdout = out
	}
	err := c.ErrorRun()
	return out.String(), err
}
