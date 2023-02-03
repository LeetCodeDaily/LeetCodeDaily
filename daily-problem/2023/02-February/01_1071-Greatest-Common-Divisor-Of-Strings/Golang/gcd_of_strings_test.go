package gcd_of_string

import (
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
	}
	for _, test := range tests {
		if got := gcdOfStrings(test.str1, test.str2); got != test.want {
			t.Errorf("gcdOfStrings(%v, %v) = %v; want %v", test.str1, test.str2, got, test.want)
		}
	}
}
