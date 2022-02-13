package sum_interpreter_go

import (
	"strings"
)

func Exec(str string) (uint32, error) {
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
