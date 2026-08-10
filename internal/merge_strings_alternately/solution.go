// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75
package mergestringsalternately

func MergeAlternately(word1 string, word2 string) string {
	minSize := min(len(word1), len(word2))
	buffer := make([]byte, 0, len(word1)+len(word2))

	for i := range minSize {
		buffer = append(buffer, word1[i])
		buffer = append(buffer, word2[i])
	}

	buffer = append(buffer, word2[minSize:]...)
	buffer = append(buffer, word1[minSize:]...)

	return string(buffer)
}
