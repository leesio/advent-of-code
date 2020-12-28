package main

import (
	"testing"
)

const (
	testCardPublicKey = 5764801
	testDoorPublicKey = 17807724
)

func TestPartOne(t *testing.T) {
	exp := 14897079
	if res := FindEncryptionKey(testCardPublicKey, testDoorPublicKey); res != exp {
		t.Errorf("Got %d for part one, expected: %d", res, exp)
	}
}
