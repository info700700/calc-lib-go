package interpreter

import (
	"strconv"
	"strings"
)

func Exec(str string) (uint32, error) {
	parts := strings.Split(str, "+")

	nums, err := convertStrsToUint32s(parts)
	if err != nil {
		return 0, err
	}

	sum := sumUint32s(nums)
	return sum, nil
}

func convertStrsToUint32s(strs []string) ([]uint32, error) {
	var nums []uint32

	for _, str := range strs {
		num, err := convertStrWithSpacesToUint32(str)
		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func convertStrWithSpacesToUint32(str string) (uint32, error) {
	trimmedStr := strings.TrimSpace(str)

	n, err := strconv.ParseUint(trimmedStr, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(n), nil
}

func sumUint32s(nums []uint32) uint32 {
	sum := uint32(0)

	for _, num := range nums {
		sum += num
	}

	return sum
}
