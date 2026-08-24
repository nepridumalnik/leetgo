// https://leetcode.com/problems/reverse-bits
package reversebits

import "math"

func reverseBitsV1(n int) int {
	out, pow := 0, 0
	for n != 0 {
		bit := (n % 2)
		n /= 2

		out += bit * (int(math.Pow(2, float64(31-pow))))
		pow++
	}

	return out
}

var ReverseBits func(int) int = reverseBitsV1
