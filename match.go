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

func Match(expr string) *Compound {
	return Do().Match(expr)
}

func (c *Compound) Match(expr string) *Compound {
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
			if r.MatchString(line) {
				_, err := fmt.Fprintln(c.Stdout, line)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func Replace(expr, repl string) *Compound {
	return Do().Replace(expr, repl)
}

func (c *Compound) Replace(expr, repl string) *Compound {
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
			_, err := fmt.Fprintln(c.Stdout, r.ReplaceAllString(line, repl))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func MatchReplace(expr, repl string) *Compound {
	return Do().MatchReplace(expr, repl)
}

func (c *Compound) MatchReplace(expr, repl string) *Compound {
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
			if r.MatchString(line) {
				_, err := fmt.Fprintln(c.Stdout, r.ReplaceAllString(line, repl))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
