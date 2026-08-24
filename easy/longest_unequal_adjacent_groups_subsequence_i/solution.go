// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i
package longestunequaladjacentgroupssubsequencei

func getLongestSubsequenceV1(words []string, groups []int) []string {
	subsequence := make([]string, 0, len(words))
	lastGroup := -1

	for i := range words {
		if lastGroup != groups[i] {
			lastGroup = groups[i]
			subsequence = append(subsequence, words[i])
		}
	}

	return subsequence
}

var GetLongestSubsequence func([]string, []int) []string = getLongestSubsequenceV1
