package saber

import (
	"net/http"
	"os"
	"time"
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
	HttpClient  *http.Client
	Debug       bool
}

func New() *Script {
	return &Script{
		ExitOnError: true,
		HttpClient: &http.Client{
			Timeout: time.Second * 5,
		},
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

func (s *Script) Log(layer int, a ...interface{}) {
	if s.Debug {
		a = append([]interface{}{strings.Repeat("+", layer+1)}, a...)
		fmt.Println(a...)
	}
}
