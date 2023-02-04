package check_inclusion

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}

	for _, test := range tests {
		if got := checkInclusion(test.s1, test.s2); got != test.want {
			t.Errorf("checkInclusion(%q, %q) = %v", test.s1, test.s2, got)
		}
	}
}
