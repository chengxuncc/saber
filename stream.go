package saber

import (
	"bufio"
	"fmt"
	"io/ioutil"
)

type LineTransform func(line string) (newLine string, ok bool)

func (c *Compound) LineInit(init func(cmd *Command) (LineTransform, error)) *Compound {
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

func (c *Compound) Line(transform LineTransform) *Compound {
	return c.LineInit(func(cmd *Command) (LineTransform, error) {
		return transform, nil
	})
}

func (c *Compound) StringTransform(f func(string) (string, error)) *Compound {
	return c.Next(func(c *Command) error {
		b, err := ioutil.ReadAll(c.GetStdin())
		if err != nil {
			return err
		}
		text, err := f(string(b))
		if err != nil {
			return err
		}
		_, err = c.Stdout.Write([]byte(text))
		return err
	})
}
