// +build gofuzzbeta

package fuzz

import (
	"strings"
	"testing"
)

func FuzzSample(f *testing.F) {
	f.Add(uint(1), "x", []byte{10, 20})
	f.Fuzz(func(t *testing.T, i1 uint, i2 string, i3 []uint8) {
		if len(i2) == 0 || len(i3) == 0 || i1 == 0 {
			t.Skip()
		}
		l := strings.Repeat(i2, int(i1))
		r := strings.Repeat(string(i3), int(i1))
		if l == r {
			t.Errorf("%s == %s (%v, %v, %v)", l, r, i1, i2, i3)
		}
	})
}
