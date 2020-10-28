package saber

import (
	"bufio"
	"fmt"
	"regexp"
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

func Regex(expr string, replace ...string) *Compound {
	return Do().Regex(expr, replace...)
}

func (c *Compound) Regex(expr string, replace ...string) *Compound {
	var repl string
	if len(replace) > 0 {
		repl = replace[0]
	}
	return c.Call(func(c *Command) error {
		if c.Stdin == nil {
			return nil
		}
		r, err := regexp.Compile(expr)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(c.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if repl != "" {
				_, err := fmt.Fprintln(c.Stdout, r.ReplaceAllString(line, repl))
				if err != nil {
					return err
				}
			} else {
				if r.MatchString(line) {
					_, err := fmt.Fprintln(c.Stdout, line)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
}
