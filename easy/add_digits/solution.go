// https://leetcode.com/problems/add-digits
package adddigits

func AddDigits(num int) int {
	return (num-1)%9 + 1
}
