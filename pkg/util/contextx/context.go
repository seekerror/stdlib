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
