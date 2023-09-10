package iox

import "sync/atomic"

// AsyncCloser is a guarded quit chan. It is safer and more convenient by allowing
// multiple calls to Close.
type AsyncCloser interface {
	// IsClosed returns true iff closed.
	IsClosed() bool
	// Closed returns the closer quit chan.
	Closed() <-chan struct{}
	// Close closes the quit chan. Idempotent.
	Close()
}

type closer struct {
	quit   chan struct{}
	closed atomic.Bool
}

func NewAsyncCloser() AsyncCloser {
	return &closer{quit: make(chan struct{})}
}

func (c *closer) IsClosed() bool {
	return c.closed.Load()
}

func (c *closer) Closed() <-chan struct{} {
	return c.quit
}

func (c *closer) Close() {
	if c.closed.CompareAndSwap(false, true) {
		close(c.quit)
	}
}
