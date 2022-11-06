package ikkle

import "testing"

const fail = "\033[31m got %v, want %v \033[0m"

func Eq[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf(fail, got, want)
	}
}
