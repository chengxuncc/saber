package saber

import (
	"os"
)

var (
	Main = New()
)

type Script struct {
	ExitOnError bool
	ExitCode    int
	Error       error
	NullStdout  bool
	NullStderr  bool
}

func New() *Script {
	return &Script{
		ExitOnError: true,
	}
}

func Do() *Compound {
	return Main.Do()
}

func (s *Script) Do() *Compound {
	c := &Compound{
		Script:   s,
		Commands: make([]*Command, 0, 3),
	}
	if !s.NullStdout {
		c.Stdout = os.Stdout
	}
	if !s.NullStderr {
		c.Stderr = os.Stderr
	}
	return c
}

func Run(comps ...*Compound) {
	Main.Run(comps...)
}

func (s *Script) Run(comps ...*Compound) {
	for _, comp := range comps {
		comp.Script = s
		comp.Run()
	}
}
