package go_1_18

import (
	"testing"
)

func IsEqual(a, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func FuzzTesting(f *testing.F) {
	f.Add([]byte{10, 11, 13}, []byte{10, 12, 13})
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		IsEqual(a, b)
	})
}
