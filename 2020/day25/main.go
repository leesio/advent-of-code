package main

import (
	"fmt"
)

const (
	doorPublicKey = 8987316
	cardPublicKey = 14681524
)

func main() {
	fmt.Printf("part 1: %d\n", FindEncryptionKey(cardPublicKey, doorPublicKey))
}

func FindEncryptionKey(cardPK, doorPK int) int {
	doorLoopSize := FindLoopSize(7, doorPK)
	return TransformSubjectNumber(cardPK, doorLoopSize)
}
func FindLoopSize(s, target int) int {
	val := 1
	for l := 1; l < 10_000_000; l++ {
		val = (val * s) % 20201227
		if val == target {
			return l
		}
	}
	panic(fmt.Errorf("failed to find loop size"))
}

func TransformSubjectNumber(s, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val = (val * s) % 20201227
	}
	return val
}
