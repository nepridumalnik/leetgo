// https://leetcode.com/problems/number-of-1-bits
package numberof1bits

func hammingWeightV1(n int) int {
	sum := 0
	for n != 0 {
		bit := (n % 2)
		n /= 2
		sum += bit
	}

	return sum
}

var HammingWeight func(int) int = hammingWeightV1
