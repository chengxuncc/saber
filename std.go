package saber

import (
	"io"
	"os"
)

type Std struct {
	Parent   *Std
	Stdin    io.ReadCloser
	Stdout   io.WriteCloser
	Stderr   io.WriteCloser
	Combined bool
}

func (s *Std) GetStdin() io.ReadCloser {
	for ptr := s; ptr != nil; ptr = ptr.Parent {
		if ptr.Stdin != nil {
			return ptr.Stdin
		}
	}
	return nil
}

func (s *Std) GetStdout() io.WriteCloser {
	for ptr := s; ptr != nil; ptr = ptr.Parent {
		if ptr.Stdout != nil {
			return ptr.Stdout
		}
	}
	return nil
}

func (s *Std) GetStderr() io.WriteCloser {
	if s.Combined {
		return s.GetStdout()
	}
	for ptr := s; ptr != nil; ptr = ptr.Parent {
		if ptr.Stderr != nil {
			return ptr.Stderr
		}
	}
	return nil
}

func (s *Std) SetStdin(stream io.ReadCloser) (err error) {
	if s == nil {
		return
	}
	if s.Stdin != nil && s.Stdin != os.Stdin {
		err = s.Stdin.Close()
		if err != nil {
			return
		}
	}
	s.Stdin = stream
	return
}

func (s *Std) SetStdout(stream io.WriteCloser) (err error) {
	if s == nil {
		return
	}
	if s.Stdout != nil && s.Stdout != os.Stdout {
		err = s.Stdout.Close()
		if err != nil {
			return
		}
	}
	s.Stdout = stream
	return
}

func (s *Std) SetStderr(stream io.WriteCloser) (err error) {
	if s == nil {
		return
	}
	if s.Stderr != nil && s.Stderr != os.Stderr {
		err = s.Stderr.Close()
		if err != nil {
			return
		}
	}
	s.Stderr = stream
	return
}
