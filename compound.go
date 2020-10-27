package saber

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type Command func(c *Compound) error

type Compound struct {
	Script   *Script
	Stdin    io.Reader
	Stdout   io.Writer
	Stderr   io.Writer
	Commands []Command
}

func Do() *Compound {
	return &Compound{
		Commands: make([]Command, 0, 3),
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
	if c.Stdout == nil {
		c.Stdout = os.Stdout
	}
	for i := 0; i < len(c.Commands); i++ {
		cmd := c.Commands[i]
		err := cmd(c)
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
	if c.Stdout != nil {
		return "", errors.New("saber: Stdout already set")
	}
	out := &bytes.Buffer{}
	c.Stdout = out
	err := c.ErrorRun()
	return out.String(), err
}
