package sum_interpreter_go

import (
	"strconv"
	"strings"
)

func convertStrsWithSpacesToUint32s(strs []string) ([]uint32, error) {
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

	n, err := parseUint32(trimmedStr)
	if err != nil {
		return 0, err
	}

	return uint32(n), nil
}

func parseUint32(str string) (uint32, error) {
	n, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(n), nil
}
