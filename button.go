package debounce

import (
	"sync"
	"time"
)

// Button which can be debounced. The normallyClosed variable determines whether it's normally open, or normally
// closed. The onClick function is executed depending on the debounce rate.
func Button(onClick func(), normallyClosed bool) *Switch {
	return &Switch{
		m:          &sync.Mutex{},
		now:        time.Now,
		onClick:    onClick,
		closed:     normallyClosed,
		bounceTime: time.Millisecond * 10,
	}
}

// Switch that can be open or closed by default.
type Switch struct {
	m          *sync.Mutex
	onClick    func()
	closed     bool
	now        func() time.Time
	lastChange time.Time
	bounceTime time.Duration
}

// SetState is used to receive a sample from the GPIO. It may be noisy, so the rate of bounces is limited
// to 1 per the bounce time of 10ms.
func (s *Switch) SetState(closed bool) {
	s.m.Lock()
	defer s.m.Unlock()
	n := s.now()
	if s.closed != closed && s.lastChange.Before(n.Add(-s.bounceTime)) {
		s.lastChange = n
		s.onClick()
	}
}
