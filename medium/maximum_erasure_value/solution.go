// https://leetcode.com/problems/maximum-erasure-value
package maximumerasurevalue

func maximumUniqueSubarrayV1(nums []int) int {
	indices := map[int]int{}
	mapSum, left, longest := 0, 0, 0

	for i, n := range nums {
		if idx, ok := indices[n]; ok {
			for left <= idx {
				delete(indices, nums[left])
				mapSum -= nums[left]
				left++
			}
		}

		indices[n] = i
		mapSum += nums[i]

		longest = max(longest, mapSum)
	}

	return longest
}

var MaximumUniqueSubarray func(nums []int) int = maximumUniqueSubarrayV1
