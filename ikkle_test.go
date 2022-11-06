package ikkle

import "testing"

func TestIkkle(t *testing.T) {

	t.Run("int", func(t *testing.T) {
		Eq(t, 1, 1)
	})

	t.Run("string", func(t *testing.T) {
		Eq(t, "hello", "hello")
	})

	t.Run("float", func(t *testing.T) {
		Eq(t, 1.234, 1.234)
	})
}
