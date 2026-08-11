// https://leetcode.com/problems/missing-number
package excelsheetcolumntitle

func MissingNumber(nums []int) int {
	n := len(nums)
	sum := (n * (n + 1)) / 2

	for _, v := range nums {
		sum -= v
	}

	return sum
}
