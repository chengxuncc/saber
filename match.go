package saber

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chengxuncc/saber/internal/x"
)

type StringTransform func(line string) (newLine string, ok bool)

func (c *Compound) StreamInit(init func(cmd *Command) (StringTransform, error)) *Compound {
	return c.Next(func(c *Command) error {
		if c.Stdin == nil {
			return nil
		}
		transform, err := init(c)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(c.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			newLine, ok := transform(line)
			if ok {
				_, err := fmt.Fprintln(c.Stdout, newLine)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (c *Compound) Stream(transform StringTransform) *Compound {
	return c.StreamInit(func(cmd *Command) (StringTransform, error) {
		return transform, nil
	})
}

func (c *Compound) Grep(sub string) *Compound {
	return c.Stream(func(line string) (newLine string, ok bool) {
		ok = strings.Contains(line, sub)
		if ok {
			newLine = line
		}
		return
	})
}

func (c *Compound) Match(expr string) *Compound {
	return c.StreamInit(func(cmd *Command) (StringTransform, error) {
		r, err := regexp.Compile(expr)
		if err != nil {
			return nil, err
		}
		return func(line string) (newLine string, ok bool) {
			ok = r.MatchString(line)
			if ok {
				newLine = line
			}
			return
		}, nil
	})
}

func (c *Compound) Replace(expr, repl string) *Compound {
	return c.StreamInit(func(cmd *Command) (StringTransform, error) {
		r, err := regexp.Compile(expr)
		if err != nil {
			return nil, err
		}
		return func(line string) (string, bool) {
			return r.ReplaceAllString(line, repl), true
		}, nil
	})
}

func (c *Compound) MatchReplace(expr, repl string) *Compound {
	return c.StreamInit(func(cmd *Command) (StringTransform, error) {
		r, err := regexp.Compile(expr)
		if err != nil {
			return nil, err
		}
		return func(line string) (string, bool) {
			if r.MatchString(line) {
				return r.ReplaceAllString(line, repl), true
			} else {
				return "", false
			}
		}, nil
	})
}

func (c *Compound) ReplaceString(old, new string) *Compound {
	return c.Stream(func(line string) (string, bool) {
		return strings.ReplaceAll(line, old, new), true
	})
}

func ReplaceFile(expr, repl, file string) *Compound {
	return Do().ReplaceFile(expr, repl, file)
}

func (c *Compound) ReplaceFile(expr, repl, file string) *Compound {
	return c.Next(func(c *Command) error {
		r, err := regexp.Compile(expr)
		if err != nil {
			return err
		}
		src, err := os.Open(file)
		if err != nil {
			return err
		}
		defer src.Close()

		tempFile := filepath.Join(os.TempDir(), "saber-"+x.RandString(8)+"-"+path.Base(file))
		dst, err := os.OpenFile(tempFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer dst.Close()

		scanner := bufio.NewScanner(src)
		for scanner.Scan() {
			line := scanner.Text()
			_, err := fmt.Fprintln(dst, r.ReplaceAllString(line, repl))
			if err != nil {
				return err
			}
		}
		src.Close()
		dst.Close()
		return Mv(tempFile, file).ErrorRun()
	})
}
