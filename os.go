package saber

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/chengxuncc/saber/internal/x"
)

func Pwd() *Compound {
	return Do().Pwd()
}

func (c *Compound) Pwd() *Compound {
	return c.Call(func(c *Command) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(c.Stdout, wd)
		return err
	})
}

func Cat(file string) *Compound {
	return Do().Cat(file)
}

func (c *Compound) Cat(file string) *Compound {
	return c.Call(func(c *Command) error {
		if c.Stdin != nil {
			return errors.New("saber: Stdin is already set")
		}
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(c.Stdout, f)
		if err != nil {
			return err
		}
		return nil
	})
}

func Cd(dir string) *Compound {
	return Do().Cd(dir)
}

func (c *Compound) Cd(dir string) *Compound {
	return c.Call(func(c *Command) error {
		return os.Chdir(dir)
	})
}

func To(file string) *Compound {
	return Do().To(file)
}

func (c *Compound) To(file string) *Compound {
	return c.Call(func(c *Command) error {
		if c.Stdin == nil {
			c.Stdin = &x.Buffer{}
		}
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, c.Stdin)
		if err != nil {
			return err
		}
		return nil
	})
}

func Append(file string) *Compound {
	return Do().Append(file)
}

func (c *Compound) Append(file string) *Compound {
	return c.Call(func(c *Command) error {
		if c.Stdin == nil {
			c.Stdin = &x.Buffer{}
		}
		f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, c.Stdin)
		if err != nil {
			return err
		}
		return nil
	})
}

func Mv(oldpath, newpath string) *Compound {
	return Do().Mv(oldpath, newpath)
}

func (c *Compound) Mv(oldpath, newpath string) *Compound {
	return c.Call(func(c *Command) error {
		err := os.Rename(oldpath, newpath)
		if err == nil {
			return nil
		}
		src, err := os.Open(oldpath)
		if err != nil {
			return err
		}
		defer src.Close()
		dst, err := os.OpenFile(newpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
		_ = os.Remove(oldpath)
		return nil
	})
}

func Cp(oldpath, newpath string) *Compound {
	return Do().Cp(oldpath, newpath)
}

func (c *Compound) Cp(oldpath, newpath string) *Compound {
	return c.Call(func(c *Command) error {
		src, err := os.Open(oldpath)
		if err != nil {
			return err
		}
		defer src.Close()
		dst, err := os.OpenFile(newpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
		return nil
	})
}
