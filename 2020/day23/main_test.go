package main

import (
	"testing"
)

var testInput = "389125467"

func TestPartOne(t *testing.T) {
	exp := "67384529"
	if res := PartOne(testInput); res != exp {
		t.Errorf("Got %s, exp: %s", res, exp)
	}
}

func BenchmarkMoving(b *testing.B) {
	r := NewRing(testInput)
	r.Pad()
	for i := 0; i < b.N; i++ {
		r.Move()
	}
}
