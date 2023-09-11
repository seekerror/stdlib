package contextx

import "context"

// IsCancelled returns true iff the context is cancelled.
func IsCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// WithQuitCancel creates a cancellable context, which is also closed if the quit chan closes.
func WithQuitCancel(ctx context.Context, quit <-chan struct{}) (context.Context, context.CancelFunc) {
	wctx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-quit:
			cancel()
		case <-wctx.Done():
		}
	}()

	return wctx, cancel
}
