package aaa

import(
	"fmt"
	"testing"
)

func gusu(s int) bool {
	if s%2 == 0 {
		return true
	} else {
		return false
	}
}

func TestFlagParser(t *testing.T) {
	var flagtests = []struct {
		in  int
		out bool
	}{
		{2, true}, {3, false}, {128, true},
		{1351, false}, {1958, true},
	}

	for _, tt := range flagtests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			result := gusu(tt.in)
			if result != tt.out {
				t.Errorf("tt.in: %d, tt.out: %t, result: %t", tt.in, tt.out, result)
			}
		})
	}
}
