// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero
package findnuniqueintegerssumuptozero

func sumZeroV1(n int) []int {
	out := make([]int, n)
	center := n / 2

	if n%2 == 1 {
		for i := 0; i <= center; i++ {
			out[center-i] = i
			out[center+i] = -i
		}
	} else {
		for i := 1; i <= center; i++ {
			out[i-1] = i
			out[n-i] = -i
		}
	}

	return out
}

func sumZeroV2(n int) []int {
	out := make([]int, n)
	center := n / 2

	for i := 1; i <= center; i++ {
		out[i-1] = -i
		out[n-i] = i
	}

	return out
}

var SumZero func(n int) []int = sumZeroV2
