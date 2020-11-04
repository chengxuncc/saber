package saber

type CallFunc func(cmd *Command) error

type Command struct {
	Std
	Compound  *Compound
	CallQueue []CallFunc
}

func (c *Command) Queue(fn CallFunc) *Command {
	if fn == nil {
		return c
	}
	c.CallQueue = append(c.CallQueue, fn)
	return c
}
