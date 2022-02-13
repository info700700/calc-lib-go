package sum_interpreter_go

import (
	"errors"
	"strings"
)

var ErrEmptyExp = errors.New("empty exp")

func Exec(str string) (uint32, error) {
	if len(str) == 0 {
		return 0, ErrEmptyExp
	}

	parts := strings.Split(str, "+")

	nums, err := convertStrsWithSpacesToUint32s(parts)
	if err != nil {
		return 0, err
	}

	sum := sumUint32s(nums)
	return sum, nil
}

func sumUint32s(nums []uint32) uint32 {
	sum := uint32(0)

	for _, num := range nums {
		sum += num
	}

	return sum
}
