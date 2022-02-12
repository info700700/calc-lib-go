package interpreter_test

import (
	"testing"

	"github.com/info700700/calc-lib-go/interpreter"
)

func TestExec(t *testing.T) {
	testCases := []struct {
		str      string
		expected uint32
	}{
		{"1", 1},
		{" 1", 1},  // Пробел в начале строки.
		{"1 ", 1},  // Пробел в конце строки.
		{" 1 ", 1}, // Пробел в начале и в конце строки.
		{"1+1", 2},
		{"1 + 1", 2},
		{"1+1+1", 3},
		{"10+20", 30},
	}

	for _, tc := range testCases {
		actual, err := interpreter.Exec(tc.str)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if actual != tc.expected {
			t.Errorf("Exec(%q) == %d, expected %d", tc.str, actual, tc.expected)
		}
	}
}
