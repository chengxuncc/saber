package x

import (
	"bytes"
	"io"
)

type ReadCloser interface {
	CloseRead() error
}

func Close(i interface{}) error {
	var errs []error
	if closer, ok := i.(ReadCloser); ok {
		errs = append(errs, closer.CloseRead())
	}
	closer, ok := i.(io.Closer)
	if ok {
		errs = append(errs, closer.Close())
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

type Buffer struct {
	bytes.Buffer
}

func (b *Buffer) Close() error { return nil }
