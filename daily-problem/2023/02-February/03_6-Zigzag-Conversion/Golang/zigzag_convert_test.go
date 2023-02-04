package zigzag_convert

import (
	"testing"
)

func TestConvert(t *testing.T) {
	testcases := []struct {
		s      string
		numRow int
		want   string
	}{
		{
			s:      "PAYPALISHIRING",
			numRow: 3,
			want:   "PAHNAPLSIIGYIR",
		},
		{
			s:      "PAYPALISHIRING",
			numRow: 4,
			want:   "PINALSIGYAHRPI",
		},
		{
			s:      "A",
			numRow: 1,
			want:   "A",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.s, func(t *testing.T) {
			if got := convert(tc.s, tc.numRow); got != tc.want {
				t.Errorf("convert() = %v, want %v", got, tc.want)
			}
		})
	}
}
