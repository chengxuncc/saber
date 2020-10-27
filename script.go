package saber

var (
	Main = New()
)

type Script struct {
	ExitOnError   bool
	ReturnCode    int
	Error         error
	DisableOutput bool
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

func (s *Script) Do() *Compound {
	comp := Do()
	comp.Script = s
	return comp
}

func (s *Script) Run(comps ...*Compound) {
	for _, comp := range comps {
		comp.Script = s
		comp.Run()
	}
}
