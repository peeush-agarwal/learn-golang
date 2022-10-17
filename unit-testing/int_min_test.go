package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	got := IntMin(2, -2)
	want := -2

	if got != want {
		fmt.Printf("Got %q, but expected %q\n", got, want)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{0, -1, -1},
		{2, -2, -2},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d,%d", test.a, test.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(test.a, test.b)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}

		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
