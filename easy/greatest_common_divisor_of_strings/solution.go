// https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75
package greatestcommondivisorofstrings

func isValid(substr, str string) bool {
	if len(str)%len(substr) != 0 {
		return false
	}

	for i := range len(str) / len(substr) {
		for j := range len(substr) {
			if str[(i*len(substr))+j] != substr[j] {
				return false
			}
		}
	}

	return true
}

func GcdOfStrings(str1 string, str2 string) string {
	result := ""
	for i := 1; i <= min(len(str1), len(str2)); i++ {
		tmp := str1[0:i]
		if isValid(tmp, str1) && isValid(tmp, str2) {
			if len(result) < len(tmp) {
				result = tmp
			}
		}
		continue
	}

	return result
}
