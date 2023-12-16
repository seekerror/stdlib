package iox

import (
	"context"
	"sync/atomic"
)

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

func NewAsyncCloserWithCancel(ctx context.Context) AsyncCloser {
	return WithCancel(ctx, NewAsyncCloser())
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

// WithCancel closes the closer if the context is cancelled. Returns the closer for convenience.
func WithCancel(ctx context.Context, closer AsyncCloser) AsyncCloser {
	return WithQuit(ctx.Done(), closer)
}

// WithQuit closes the closer if the quit chan closes. Returns the closer for convenience.
func WithQuit(quit <-chan struct{}, closer AsyncCloser) AsyncCloser {
	go func() {
		select {
		case <-quit:
			closer.Close()
		case <-closer.Closed():
		}
	}()

	return closer
}
