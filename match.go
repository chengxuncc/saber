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
		transform, err := init(c)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(c.GetStdin())
		for scanner.Scan() {
			line := scanner.Text()
			newLine, ok := transform(line)
			if ok {
				_, err := fmt.Fprintln(c.GetStdout(), newLine)
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

func (c *Compound) MatchDelete(expr string) *Compound {
	return c.StreamInit(func(cmd *Command) (StringTransform, error) {
		r, err := regexp.Compile(expr)
		if err != nil {
			return nil, err
		}
		return func(line string) (newLine string, ok bool) {
			ok = !r.MatchString(line)
			newLine = line
			return
		}, nil
	})
}

func (c *Compound) ReplaceString(old, new string) *Compound {
	return c.Stream(func(line string) (string, bool) {
		return strings.ReplaceAll(line, old, new), true
	})
}

func ReplaceFile(expr, repl string, file ...string) *Compound {
	return Do().ReplaceFile(expr, repl, file...)
}

func (c *Compound) ReplaceFile(expr, repl string, file ...string) *Compound {
	return c.Next(func(c *Command) error {
		r, err := regexp.Compile(expr)
		if err != nil {
			return err
		}
		for _, f := range file {
			src, err := os.Open(f)
			if err != nil {
				return err
			}

			tempFile := filepath.Join(os.TempDir(), "saber-"+x.RandString(8)+"-"+path.Base(f))
			dst, err := os.OpenFile(tempFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
			if err != nil {
				src.Close()
				return err
			}
			scanner := bufio.NewScanner(src)
			for scanner.Scan() {
				line := scanner.Text()
				_, err := fmt.Fprintln(dst, r.ReplaceAllString(line, repl))
				if err != nil {
					src.Close()
					dst.Close()
					return err
				}
			}
			src.Close()
			dst.Close()
			err = Mv(tempFile, f).ErrorRun()
			if err != nil {
				return err
			}
		}
		return nil
	})
}
