// https://leetcode.com/problems/happy-number
package happynumber

func IsHappy(n int) bool {
	count := 0
	for count < 100 {
		count++

		tmp := 0

		for n != 0 {
			num := (n % 10)
			n /= 10
			tmp += num * num
		}

		if tmp == 1 {
			return true
		}

		n = tmp
	}

	return false
}
