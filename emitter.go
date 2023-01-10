package was

// Emitter allows to register handlers for events and emit events.

type Emitter[T any] struct {
	handlers map[string][]func(T)
}

// NewEmitter creates a new Emitter.
func NewEmitter[T any]() *Emitter[T] {
	return &Emitter[T]{
		handlers: make(map[string][]func(T)),
	}
}

// On registers a handler for the given event.
func (e *Emitter[T]) On(event string, handler func(T)) {
	e.handlers[event] = append(e.handlers[event], handler)
}

// Emit emits the given event with the given data.
func (e *Emitter[T]) Emit(event string, data T) {
	for _, handler := range e.handlers[event] {
		handler(data)
	}
}

// Off removes all handlers for the given event.
func (e *Emitter[T]) Off(event string) {
	delete(e.handlers, event)
}
