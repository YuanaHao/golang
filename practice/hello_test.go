package main

import "testing"

func testhello(t *testing.T) {
	got := hello()
	want := "hello,world"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}

}
