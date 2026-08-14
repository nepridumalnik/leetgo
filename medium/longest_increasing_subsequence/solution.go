// https://leetcode.com/problems/longest-increasing-subsequence
package longestincreasingsubsequence

import "slices"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sub := make([]int, 0, len(nums))
	sub = append(sub, nums[0])

	for _, n := range nums[1:] {
		switch {
		case sub[len(sub)-1] < n:
			sub = append(sub, n)
		default:
			idx, _ := slices.BinarySearch(sub, n)
			sub[idx] = n
		}
	}

	return len(sub)
}

var LengthOfLIS func([]int) int = lengthOfLIS
