package signalx

import (
	"os"
	"os/signal"
)

// TrapInterrupt returns a quit channel that is closed if SIGINT or SIGKILL.
func TrapInterrupt() <-chan struct{} {
	quit := make(chan struct{})
	go func() {
		defer close(quit)

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill)
		<-ch
	}()
	return quit
}
