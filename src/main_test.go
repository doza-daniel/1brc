package main

import (
	"testing"
)

func Test_mustParseFloat64(t *testing.T) {
	tt := []struct {
		in     string
		expect float64
	}{
		{
			in:     "12.3",
			expect: 12.3,
		},
		{
			in:     "-12.3",
			expect: -12.3,
		},
		{
			in:     "1.3",
			expect: 1.3,
		},
		{
			in:     "-1.3",
			expect: -1.3,
		},
		{
			in:     "0.0",
			expect: 0.0,
		},
		{
			in:     "-0.0",
			expect: 0.0,
		},
		{
			in:     "1.0",
			expect: 1.0,
		},
		{
			in:     "-1.0",
			expect: -1.0,
		},
		{
			in:     "0.1",
			expect: 0.1,
		},
		{
			in:     "-0.1",
			expect: -0.1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			got := mustParseFloat64(tc.in)
			if tc.expect != got {
				t.Errorf("expected %+v, got %+v", tc.expect, got)
			}
		})
	}
}
