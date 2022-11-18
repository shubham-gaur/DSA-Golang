package array

import (
	"fmt"
	"testing"
)

func TestValidateParenthesisStack(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			"()()()()()()()()()()(*))",
			true,
		},
		{
			"(*))()()()()()()()()()()",
			true,
		},
		{
			"(((*)))",
			true,
		},
		{
			"()())*",
			false,
		},
		{
			"()*)",
			true,
		},
		{
			"(((*)",
			false,
		},
		{
			"",
			true,
		},
		{
			"((*)",
			true,
		},
		{
			"(((*****)",
			true,
		},
	}

	for idx, tc := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testName := fmt.Sprintf("%d", idx)

		t.Run(testName, func(t *testing.T) {
			got := ValidateParenthesisStack(tc.input)

			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}
