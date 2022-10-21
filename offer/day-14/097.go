package day_14

func numDistinct(s string, t string) int {
	arr := [2][]int{}
	arr[0], arr[1] = make([]int, len(t)+1), make([]int, len(t)+1)
	arr[0][0] = 1
	for i := 0; i < len(s); i++ {
		arr[(i+1)%2][0] = 1
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				arr[(i+1)%2][j+1] = arr[i%2][j] + arr[i%2][j+1]
			} else {
				arr[(i+1)%2][j+1] = arr[i%2][j+1]
			}
		}
	}
	return arr[len(s)%2][len(t)]
}

func numDistinctV1(s string, t string) int {
	arr := make([][]int, len(s)+1)
	arr[0] = make([]int, len(t)+1)
	for i := 0; i < len(s); i++ {
		arr[(i+1)%2] = make([]int, len(t)+1)
		for j := 0; j < len(t); j++ {
			if j == 0 {
				arr[i][j] = 1
			}
			if s[i] == t[j] {
				arr[(i+1)%2][j+1] = arr[i][j] + arr[i][j+1]
			} else {
				arr[(i+1)%2][j+1] = arr[i][j+1]
			}
		}
	}
	return arr[len(s)][len(t)]
}

//   f(i,j) ==> s[i],t[j]
//   s[i] == f[j]   == > f(i,j) = f(i-1,j-1)+1 + f(i-i,j)
//   s[i] != f[j]   == > f(i,j) =  f(i-i,j)
//   f(0,j) = 0
//   f(i,0) = 1
//   f(len(s),len(t))
