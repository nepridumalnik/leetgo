// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75
package merge_strings_alternately

func MergeAlternately(word1 string, word2 string) string {
	minSize := min(len(word1), len(word2))
	buffer := make([]byte, 0, len(word1)+len(word2))
	a := []byte(word1)
	b := []byte(word2)

	for i := range minSize {
		buffer = append(buffer, byte(a[i]))
		buffer = append(buffer, byte(b[i]))
	}

	switch {
	case len(a) == len(b):
		break
	case len(a) > minSize:
		buffer = append(buffer, a[minSize:]...)
	default:
		buffer = append(buffer, b[minSize:]...)
	}

	return string(buffer)
}
