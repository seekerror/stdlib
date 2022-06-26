package iox

import "time"

// Pulse is a simple manual pulse. Signals are deduplicated, if not consumed.
type Pulse struct {
	out chan bool
}

func NewPulse() *Pulse {
	return &Pulse{out: make(chan bool, 1)}
}

// Emit signals the pulse, if not already signaled. Returns true if signaled.
func (p *Pulse) Emit() bool {
	select {
	case p.out <- true:
		return true
	default:
		return false
	}
}

// Chan returns the emit chan for consumption.
func (p *Pulse) Chan() <-chan bool {
	return p.out
}

func NewTickerPulse(duration time.Duration) *Pulse {
	return WithTicker(NewPulse(), time.NewTicker(duration).C)
}

func WithTicker(pulse *Pulse, ch <-chan time.Time) *Pulse {
	go func() {
		for range ch {
			pulse.Emit()
		}
	}()
	return pulse
}
