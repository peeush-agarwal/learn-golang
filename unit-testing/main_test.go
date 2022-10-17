package main

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{1, 2, 3},
	addTest{2, 3, 5},
	addTest{-1, -2, -3},
	addTest{-2, 3, 1},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		got := add(test.arg1, test.arg2)
		if got != test.expected {
			t.Errorf("Output %q not equal to expected %q", got, test.expected)
		}
	}
}
