// https://leetcode.com/problems/longest-substring-without-repeating-characters
package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstringV1(s string) int {
	chars := map[rune]int{}
	longest := 0
	left := 0

	for i, c := range s {
		if idx, ok := chars[c]; ok {
			longest = max(longest, len(chars))

			for left <= idx {
				delete(chars, rune(s[left]))
				left++
			}
		}
		chars[c] = i
	}

	return max(longest, len(chars))
}

func lengthOfLongestSubstringV2(s string) int {
	left, longest := 0, 0
	indices := [128]int{}

	for right, char := range s {
		if indices[char] > left {
			left = indices[char]
		}

		indices[char] = right + 1
		longest = max(longest, right-left+1)
	}

	return longest
}

var LengthOfLongestSubstring func(string) int = lengthOfLongestSubstringV2
