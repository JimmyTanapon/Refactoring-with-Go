package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	var cases = []struct {
		name  string
		input int
		want  string
	}{
		{name: "input 1 should be 1", input: 1, want: "1"},
		{name: "input 2 should be 2", input: 2, want: "2"},
		{name: "input 3 should be Fizz", input: 3, want: "Fizz"},
		{name: "input 4 should be 4", input: 4, want: "4"},
		{name: "input 5 should be Buzz", input: 5, want: "Buzz"},
		{name: "input 6 should be Fizz", input: 6, want: "Fizz"},
		{name: "input 7 should be 7", input: 7, want: "7"},
		{name: "input 8 should be 8", input: 8, want: "8"},
		{name: "input 9 should be Fizz", input: 9, want: "Fizz"},
		{name: "input 10 should be Buzz", input: 10, want: "Buzz"},
		{name: "input 12 should be Fizz", input: 12, want: "Fizz"},
		{name: "input 15 should be FizzBuzz", input: 15, want: "FizzBuzz"},
		{name: "input 30 should be FizzBuzz", input: 30, want: "FizzBuzz"},
		{name: "input 45 should be FizzBuzz", input: 45, want: "FizzBuzz"},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := FizzBuzz(test.input)

			if got != test.want {
				t.Errorf("got = %q, want %q", got, test.want)
			}
		})

	}
}
