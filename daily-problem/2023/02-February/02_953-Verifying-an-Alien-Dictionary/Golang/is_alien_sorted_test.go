package is_alien_sorted

import (
	"testing"
)

func TestIsAlienSorted(t *testing.T) {
	testCases := []struct {
		words []string
		order string
		want  bool
	}{
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	}

	for _, tc := range testCases {
		got := isAlienSorted(tc.words, tc.order)
		if got != tc.want {
			t.Errorf("isAlienSorted(%v, %v) = %v; want %v", tc.words, tc.order, got, tc.want)
		}
	}
}
