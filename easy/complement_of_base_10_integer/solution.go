// https://leetcode.com/problems/complement-of-base-10-integer
package complementofbase10integer

import "math"

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	num, pow := 0, 0
	for n != 0 {
		if (n % 2) == 0 {
			res := int(math.Pow(2, float64(pow)))
			num += res
		}

		n /= 2
		pow++
	}

	return num
}

var BitwiseComplement func(int) int = bitwiseComplement
