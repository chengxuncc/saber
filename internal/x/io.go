package x

import (
	"bytes"
	"io"
	"sync"
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

// Discard is an io.Writer on which all Write calls succeed
// without doing anything.
var Discard = devNull(0)

type devNull int

// devNull implements ReaderFrom as an optimization so io.Copy to
// ioutil.Discard can avoid doing unnecessary work.
var _ io.ReaderFrom = devNull(0)

func (devNull) Write(p []byte) (int, error) {
	return len(p), nil
}

func (devNull) WriteString(s string) (int, error) {
	return len(s), nil
}

var blackHolePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 8192)
		return &b
	},
}

func (devNull) ReadFrom(r io.Reader) (n int64, err error) {
	bufp := blackHolePool.Get().(*[]byte)
	readSize := 0
	for {
		readSize, err = r.Read(*bufp)
		n += int64(readSize)
		if err != nil {
			blackHolePool.Put(bufp)
			if err == io.EOF {
				return n, nil
			}
			return
		}
	}
}

func (devNull) Close() (err error) {
	return
}
