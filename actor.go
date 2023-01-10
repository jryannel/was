package was

import "time"

type Actor struct {
	*Emitter[any]
}

func NewActor() *Actor {
	return &Actor{
		Emitter: NewEmitter[any](),
	}
}

func Update(elapsed time.Duration) {
	// ...
}
