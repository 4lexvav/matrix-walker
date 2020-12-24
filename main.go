package main

import "fmt"

func flatten(matrix [][]int32) []int32 {
	var plain []int32
	n := len(matrix)
	m := n

	// validate matrix
	for _, item := range matrix {
		if len(item) != n {
			panic("Matrix is not even!")
		}
	}

	k := 0
	l := 0

	for k < m && l < n {
		for i := l; i < n; i++ {
			plain = append(plain, matrix[k][i])
		}
		k++

		for i := k; i < m; i++ {
			plain = append(plain, matrix[i][n-1])
		}
		n--

		if k < m {
			for i := n - 1; i >= l; i-- {
				plain = append(plain, matrix[m-1][i])
			}
			m--
		}

		if l < n {
			for i := m - 1; i >= k; i-- {
				plain = append(plain, matrix[i][l])
			}
			l++
		}
	}

	return plain
}

func main() {
	matrix := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix2 := [][]int32{
		{1, 2, 3, 1},
		{4, 5, 6, 4},
		{7, 8, 9, 7},
		{7, 8, 9, 7},
	}

	fmt.Println(flatten(matrix))
	fmt.Println(flatten(matrix2))
}
