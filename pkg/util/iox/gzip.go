package iox

import (
	"compress/gzip"
	"io"
)

type gzipWriter struct {
	w   *gzip.Writer
	out io.Closer
}

// NewGzipWriter is a gzip.Writer that also closes the underlying stream. Convenience wrapper.
func NewGzipWriter(out io.WriteCloser) io.WriteCloser {
	return &gzipWriter{
		w:   gzip.NewWriter(out),
		out: out,
	}
}

func (c *gzipWriter) Write(p []byte) (int, error) {
	return c.w.Write(p)
}

func (c *gzipWriter) Close() error {
	defer c.out.Close()
	return c.w.Close()
}
