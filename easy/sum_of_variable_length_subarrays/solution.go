// https://leetcode.com/problems/sum-of-variable-length-subarrays
package sumofvariablelengthsubarrays

func collectSum(subarray []int) int {
	sum := 0
	for _, n := range subarray {
		sum += n
	}

	return sum
}

// Naive approach
// TODO: think how to make faster than O(n)
func subarraySumV1(nums []int) int {
	sum := 0
	for i := range nums {
		start := max(0, i-nums[i])
		sum += collectSum(nums[start : i+1])
	}

	return sum
}

var SubarraySum func([]int) int = subarraySumV1
