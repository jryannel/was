package was

import "time"

type System interface {
	Update(elapsed time.Duration)
	World() *World
}
