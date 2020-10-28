package saber

import (
	"bufio"
	"fmt"
	"strings"
)

func Grep(sub string) *Compound {
	return Do().Grep(sub)
}

func (c *Compound) Grep(sub string) *Compound {
	return c.Call(func(c *Command) error {
		if c.Stdin == nil {
			return nil
		}
		scanner := bufio.NewScanner(c.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, sub) {
				_, err := fmt.Fprintln(c.Stdout, line)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
