package main

import "testing"

func Test_greet(t *testing.T) {
	got := greet()
	want := "Hi from Poland"
	if got != want {
		t.Errorf("greet() = %v, want %v", got, want)
	}
}
