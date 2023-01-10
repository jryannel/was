package was

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEmitterOn tests the On method.
func TestEmitterOn(t *testing.T) {
	e := NewEmitter[int]()
	e.On("test", func(data int) {
		assert.Equal(t, 10, data)
	})
	e.Emit("test", 10)
}
