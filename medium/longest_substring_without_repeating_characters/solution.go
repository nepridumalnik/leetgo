// https://leetcode.com/problems/longest-substring-without-repeating-characters
package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
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

var LengthOfLongestSubstring func(string) int = lengthOfLongestSubstring
