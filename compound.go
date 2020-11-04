package saber

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chengxuncc/saber/internal/x"
)

type Compound struct {
	Std
	Layer    int
	Script   *Script
	Commands []*Command
}

func (c *Compound) Next(fn CallFunc) *Compound {
	cmd := &Command{
		CallQueue: []CallFunc{fn},
		Compound:  c,
	}
	lastCmd := c.Current()
	if lastCmd == nil {
		// use compound Stdin on the first command
		cmd.Stdin = c.Stdin
	} else {
		if lastCmd.Stdout != nil {
			cmd.Stdin, lastCmd.Stdout = io.Pipe()
		}
	}
	c.Commands = append(c.Commands, cmd)
	return c
}

func (c *Compound) Run() {
	c.check(nil, c.ErrorRun())
}

func (c *Compound) ErrorRun() error {
	count := len(c.Commands)
	errs := make(chan error, count)
	for i := 0; i < count; i++ {
		go func(cmd *Command) {
			if cmd.Stdout == nil {
				cmd.Stdout = c.Stdout
			}
			if cmd.Stderr == nil {
				cmd.Stderr = c.Stderr
			}
			var err error
			defer func() {
				errs <- err
				_ = cmd.SetStdout(nil)
				_ = cmd.SetStderr(nil)
			}()
			for i := 0; i < len(cmd.CallQueue); i++ {
				err = cmd.CallQueue[i](cmd)
				if err != nil {
					return
				}
			}
		}(c.Commands[i])
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
	return c.check(c.ErrorOutput()).(string)
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

func (c *Compound) Int() int {
	return c.check(c.ErrorInt()).(int)
}

func (c *Compound) ErrorInt() (i int, err error) {
	out, err := c.ErrorOutput()
	if err != nil {
		return 0, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
			return
		}
	}()
	i = Int(out)
	return
}

func (c *Compound) String() string {
	return c.check(c.ErrorString()).(string)
}

func (c *Compound) ErrorString() (string, error) {
	out, err := c.ErrorOutput()
	out = strings.TrimSpace(out)
	return out, err
}

func (c *Compound) Current() *Command {
	length := len(c.Commands)
	if length == 0 {
		return nil
	}
	return c.Commands[len(c.Commands)-1]
}

func (c *Compound) Queue(fn CallFunc) *Compound {
	if fn == nil {
		return c
	}
	c.Current().Queue(fn)
	return c
}

func (c *Compound) check(i interface{}, err error) interface{} {
	if err != nil {
		x.Must(fmt.Fprintln(os.Stderr, err))
		if c.Script.ExitOnError {
			os.Exit(1)
		}
	}
	return i
}

func Group(comps ...*Compound) *Compound {
	return Do().Group(comps...)
}

func (c *Compound) Group(comps ...*Compound) *Compound {
	return c.Next(func(cmd *Command) error {
		c.Log("Group")
		for _, comp := range comps {
			comp.Script = c.Script
			comp.Stdin = cmd.Stdin
			comp.Stdout = cmd.Stdout
			comp.Stderr = cmd.Stderr
			err := comp.ErrorRun()
			if err != nil {
				return err
			}
		}
		return nil
	})
}
