package aaa

import (
	"fmt"
	"testing"
)

func TestGusu(t *testing.T) {
	var testcase = []struct {
		in  int
		out bool
	}{
		{2, true}, {3, false}, {0, true}, {1472897, false}, {-327, false},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			result := gusu(tt.in)
			if result != tt.out {
				t.Errorf("tt.in: %d, tt.out: %t, result: %t", tt.in, tt.out, result)
			}
		})
	}
}
