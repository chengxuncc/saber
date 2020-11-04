package saber

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chengxuncc/saber/internal/x"
)

var (
	Main = New()
)

type Script struct {
	Std
	ExitOnError bool
	ExitCode    int
	Error       error
	NullStdout  bool
	NullStderr  bool
	HttpClient  *http.Client
	Debug       bool
}

func New() *Script {
	s := &Script{
		Std: Std{
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		},
		ExitOnError: true,
		HttpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
	if s.NullStdout {
		s.Stdout = x.Discard
	}
	if s.NullStderr {
		s.Stderr = x.Discard
	}
	return s
}

func Do() *Compound {
	return Main.Do()
}

func (s *Script) Do() *Compound {
	c := &Compound{
		Std: Std{
			Parent: &s.Std,
		},
		Script:   s,
		Commands: make([]*Command, 0, 3),
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

func (s *Script) Log(layer int, name string, a ...interface{}) {
	if s.Debug {
		a = append([]interface{}{strings.Repeat("+", layer+1), name}, a...)
		fmt.Println(a...)
	}
}
