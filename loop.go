package was

import (
	"time"
)

type UpdateFunc func(elapsed time.Duration)

// Loop is the main loop of the game.
type Loop struct {
	// The time that the loop is running for.
	Elapsed   time.Duration
	SliceTime time.Duration
	Tick      time.Duration
	Update    UpdateFunc
	Done      chan bool
}

// NewLoop creates a new loop.
func NewLoop(sliceTime time.Duration, tickTime time.Duration, update UpdateFunc) *Loop {
	return &Loop{
		SliceTime: sliceTime,
		Tick:      tickTime,
		Update:    update,
		Done:      make(chan bool),
	}
}

// Run runs the loop.
func (l *Loop) Run() {
	for {
		select {
		case <-l.Done:
			return
		default:
			startTime := time.Now()
			if l.Update != nil {
				l.Update(l.Tick)
			}
			elapsedTime := time.Since(startTime)
			sleepTime := l.SliceTime - elapsedTime
			if sleepTime > 0 {
				time.Sleep(sleepTime)
			}
			l.Elapsed += l.Tick
		}
	}
}

func (l *Loop) Stop() {
	l.Done <- true
}
