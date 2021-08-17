package bitset_tests

import (
	"testing"
)

func AreEqual(t *testing.T, expected []int64, actual []int64) {
	if !areEqual(expected, actual) {
		t.Errorf("Assert failed")
	}
}

func areEqual(expected []int64, actual []int64) bool {
	i := 0
	j := 0

	for i < len(actual) && j < len(expected) {
		if actual[i] != expected[j] {
			return false
		}
		i++
		j++
	}

	if i < len(actual) {
		return false
	}

	if j < len(expected) {
		return false
	}

	return true
}
