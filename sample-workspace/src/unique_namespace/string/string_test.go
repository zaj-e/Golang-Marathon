package string
// This file won't be compiled into the package when you run go install or go build.
// However, when we run go test, the go tool will compile all the package files including the tests
// And synthesize a test harness that runs any functions beginning with test with a capital T

import "testing"

// We'll write a quick table-driven test for the reverse function

func Test(t * testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Backward","drawkcaB"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}	
}