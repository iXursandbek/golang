package main

import "testing"

func TestCalcSquare(t *testing.T) {
	got := CalcSquare(4, 4)
	want := 16.0

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
