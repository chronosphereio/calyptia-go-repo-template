package main

import (
	"testing"
)

func Test_greet(t *testing.T) {
	if want, got := "Hi!", greet(); want != got {
		t.Errorf("want: %s got: %s", want, got)
	}
}
