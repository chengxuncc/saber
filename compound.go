package saber

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Compound struct {
	Script   *Script
	Commands []*Command
}

func Do() *Compound {
	return &Compound{
		Commands: make([]*Command, 0, 3),
	}
}

func (c *Compound) Run() {
	err := c.ErrorRun()
	if err != nil {
		fmt.Println(err)
		if c.Script.ExitOnError {
			os.Exit(1)
		}
	}
}

func (c *Compound) ErrorRun() error {
	if c.Script == nil {
		c.Script = Main
	}
	// pipe
	for i := 1; i < len(c.Commands); i++ {
		lastOut := c.Commands[i-1].Stdout
		lastReader, ok := lastOut.(io.Reader)
		if ok {
			c.Commands[i].Stdin = lastReader
		}
	}
	lastCmd := c.Commands[len(c.Commands)-1]
	if lastCmd.Stdout == nil {
		lastCmd.Stdout = os.Stdout
	}
	for i := 0; i < len(c.Commands); i++ {
		cmd := c.Commands[i]
		err := cmd.Call(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compound) Output() string {
	res, err := c.ErrorOutput()
	if err != nil {
		fmt.Println(err)
		if c.Script.ExitOnError {
			os.Exit(1)
		}
	}
	return res
}

func (c *Compound) ErrorOutput() (string, error) {
	out := &bytes.Buffer{}
	c.Commands[len(c.Commands)-1].Stdout = out
	err := c.ErrorRun()
	return out.String(), err
}
