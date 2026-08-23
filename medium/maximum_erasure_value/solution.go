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

func maximumUniqueSubarrayV2(nums []int) int {
	sum, left, longest := 0, 0, 0
	indices := [10000]int{}

	for i, n := range nums {
		if idx := indices[n-1]; idx != 0 {
			for left < idx {
				indices[n-1] = 0
				sum -= nums[left]
				left++
			}
		}

		indices[n-1] = i + 1
		sum += nums[i]

		longest = max(longest, sum)
	}

	return longest
}

var MaximumUniqueSubarray func(nums []int) int = maximumUniqueSubarrayV2
