package saber

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

func Run(comps ...*Compound) {
	for _, comp := range comps {
		comp.Run()
	}
}

func Do() *Compound {
	return Main.Do()
}

func (s *Script) Do() *Compound {
	return &Compound{
		Script:   s,
		Commands: make([]*Command, 0, 3),
	}
}
