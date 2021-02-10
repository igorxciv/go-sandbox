package circuit_breaker

import (
	"errors"
	"fmt"
)

// Represents a state of circuit breaker
type State int

// Constant states of circuit breaker
const (
	Closed State = iota
	HalfOpen
	Open
)

var (
	ErrTooManyRequests = errors.New("too many requests")
	ErrOpenState = errors.New("circuit breaker is open")
)

func (s State) String() string {
	switch s {
	case Closed:
		return "closed"
	case HalfOpen:
		return "half-open"
	case Open:
		return "open"
	default:
		return fmt.Sprintf("unknown state: %d", s)
	}
}