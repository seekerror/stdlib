package iox

import (
	"bufio"
	"io"
)

type bufWriter struct {
	w   *bufio.Writer
	out io.Closer
}

// NewBufWriter is a buffered WriterCloser suitable for file I/O.
func NewBufWriter(out io.WriteCloser, size int) io.WriteCloser {
	return &bufWriter{
		w:   bufio.NewWriterSize(out, size),
		out: out,
	}
}

func (c *bufWriter) Write(p []byte) (int, error) {
	return c.w.Write(p)
}

func (c *bufWriter) Close() error {
	defer c.out.Close()
	return c.w.Flush()
}

type bufReader struct {
	r  *bufio.Reader
	in io.Closer
}

// NewBufReader is a buffered ReadCloser suitable for file I/O.
func NewBufReader(in io.ReadCloser, size int) io.ReadCloser {
	return &bufReader{
		r:  bufio.NewReaderSize(in, size),
		in: in,
	}
}

func (c *bufReader) Read(p []byte) (n int, err error) {
	return c.r.Read(p)
}

func (c *bufReader) Close() error {
	return c.in.Close()
}
