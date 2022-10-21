package day_14

import "fmt"

func minimumTotal(g [][]int) int {
	arr := make([][]int, len(g))

	for i := 0; i < len(g); i++ {
		arr[i] = make([]int, i+1)
		for j := 0; j < len(g[i]); j++ {
			if i == 0 && j == 0 {
				arr[0][0] = g[0][0]
				continue
			}
			if j == 0 {
				arr[i][0] = arr[i-1][0] + g[i][j]
				continue
			}

			if i == j {
				arr[i][j] = arr[i-1][j-1] + g[i][j]
				continue
			}
			arr[i][j] = min(arr[i-1][j], arr[i-1][j-1]) + g[i][j]
		}
	}
	result := arr[len(g)-1][0]
	for _, num := range arr[len(g)-1] {
		result = min(result, num)
	}
	fmt.Println(arr[len(g)-1])
	return result
}

// f(i,j)  i >=j
// j==0  >>  f(i,0)  == f(i-1,0) + g[i][0]
// i==j  >>  f(i,j) f(i,j) = f(i-i,j-1) + g[i][j]
// min
