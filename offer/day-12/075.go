package day_12

func relativeSortArray(arr1 []int, arr2 []int) []int {

	var count [1001]int
	var r []int
	for _, num := range arr1 {
		count[num]++
	}

	// print arr2 sort
	for _, num := range arr2 {
		for count[num] > 0 {
			r = append(r, num)
			count[num]--
		}
	}

	// print arr2 sort
	for i := 0; i <= 1000; i++ {
		for count[i] > 0 {
			r = append(r, i)
			count[i]--
		}
	}

	return r
}
