// https://leetcode.com/problems/pascals-triangle
package pascalstriangle

func newTriangle(numRows int) [][]int {
	triangle := make([][]int, 0, numRows)
	triangle = append(triangle, []int{1}, []int{1, 1})

	for i := 2; i < numRows; i++ {
		row := make([]int, 0, i+1)
		row = append(row, 1)

		for j := range cap(row) - 2 {
			left := triangle[i-1][j]
			right := triangle[i-1][j+1]
			num := left + right
			row = append(row, num)
		}

		row = append(row, 1)

		triangle = append(triangle, row)
	}

	return triangle
}

func generate(numRows int) [][]int {
	// Without factorial
	switch numRows {
	case 1:
		return [][]int{{1}}
	case 2:
		return [][]int{{1}, {1, 1}}
	default:
		return newTriangle(numRows)
	}
}

var Generate func(int) [][]int = generate
