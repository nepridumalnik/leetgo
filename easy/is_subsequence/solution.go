// https://leetcode.com/problems/is-subsequence
package issubsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	idx := 0
	for _, c := range t {
		if rune(s[idx]) == c {
			if idx++; idx == len(s) {
				return true
			}
		}
	}

	return false
}

var IsSubsequence func(string, string) bool = isSubsequence
