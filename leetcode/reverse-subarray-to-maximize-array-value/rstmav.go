package reverse_subarray_to_maximize_array_value

import "math"

// https://leetcode.com/problems/reverse-subarray-to-maximize-array-value/description/

const (
	MIN = math.MinInt
	MAX = math.MaxInt
)

func maxValueAfterReverse(nums []int) int {
	length := len(nums)

	result := 0
	for i := 0; i < length-1; i++ {
		result += abs(nums[i], nums[i+1])
	}

	min := MIN
	max := MAX

	for i := 0; i < length-1; i++ {
		maxDiff, minDiff := compare(nums[i], nums[i+1])
		_, min = compare(min, minDiff)
		max, _ = compare(max, maxDiff)
	}

	diff, _ := compare(0, (min-max)*2)
	for i := 0; i < length-1; i++ {
		diff, _ = compare(diff, abs(nums[0], nums[i+1])-abs(nums[i], nums[i+1]))
	}

	for i := 0; i < length-1; i++ {
		diff, _ = compare(diff, abs(nums[length-1], nums[i])-abs(nums[i+1], nums[i]))
	}
	return result + diff
}
func compare(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func abs(a, b int) int {
	max, min := compare(a, b)
	return max - min
}
