package sum_interpreter_go_test

import (
	"testing"

	interp "github.com/info700700/sum_interpreter_go"
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
		actual, err := interp.Exec(tc.str)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if actual != tc.expected {
			t.Errorf("Exec(%q) == %d, expected %d", tc.str, actual, tc.expected)
		}
	}
}

func TestExecEmptyExp(t *testing.T) {
	expStr := ""

	_, err := interp.Exec(expStr)
	if err != interp.ErrEmptyExp {
		t.Errorf("Err: %v. Expected ErrEmptyExp.", err)
	}
}

func TestExecInvalidExp(t *testing.T) {
	expStr := "1 +"

	_, err := interp.Exec(expStr)
	if err != interp.ErrInvalidExp {
		t.Errorf("Err: %v. Expected ErrInvalidExp.", err)
	}
}
