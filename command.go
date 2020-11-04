package saber

type CallFunc func(cmd *Command) error

type Command struct {
	Std
	Compound  *Compound
	CallStack []CallFunc
}

func (c *Command) Stack(fn CallFunc) *Command {
	if fn == nil {
		return c
	}
	c.CallStack = append(c.CallStack, fn)
	return c
}
