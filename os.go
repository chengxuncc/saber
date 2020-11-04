package saber

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Pwd() *Compound {
	return Do().Pwd()
}

func (c *Compound) Pwd() *Compound {
	return c.Next(func(cmd *Command) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(cmd.Stdout, wd)
		return err
	})
}

func Cat(file string) *Compound {
	return Do().Cat(file)
}

func (c *Compound) Cat(file string) *Compound {
	return c.Next(func(cmd *Command) error {
		if cmd.Stdin != nil {
			return errors.New("saber: Stdin is already set")
		}
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(cmd.Stdout, f)
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
	return c.Next(func(cmd *Command) error {
		return os.Chdir(dir)
	})
}

func Mv(oldpath, newpath string) *Compound {
	return Do().Mv(oldpath, newpath)
}

func (c *Compound) Mv(oldpath, newpath string) *Compound {
	return c.Next(func(cmd *Command) error {
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
	return c.Next(func(cmd *Command) error {
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

func (c *Compound) In(file string) *Compound {
	return c.Queue(func(cmd *Command) error {
		err := cmd.SetStdin(nil)
		if err != nil {
			return err
		}
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		cmd.Stdout = f
		return nil
	})
}

func (c *Compound) To(file string) *Compound {
	return c.Queue(func(cmd *Command) error {
		err := cmd.SetStdout(nil)
		if err != nil {
			return err
		}
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		cmd.Stdout = f
		return nil
	})
}

func (c *Compound) App(file string) *Compound {
	return c.Queue(func(cmd *Command) error {
		err := cmd.SetStdout(nil)
		if err != nil {
			return err
		}
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		cmd.Stdout = f
		return nil
	})
}

func (c *Compound) Err(file string) *Compound {
	return c.Queue(func(cmd *Command) error {
		err := cmd.SetStderr(nil)
		if err != nil {
			return err
		}
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		cmd.Stderr = f
		return nil
	})
}
