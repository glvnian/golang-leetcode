package day_02

func twoSum(numbers []int, target int) []int {
	var m = make(map[int]int)
	for i, num := range numbers {
		if _, ok := m[target-num]; !ok {
			m[num] = i
			continue
		}
		return []int{m[target-num], i}
	}
	return nil
}

func twoSumV1(numbers []int, target int) []int {

	var i = 0
	var j = len(numbers) - 1

	for i < j && i < len(numbers) {

		if target == numbers[i]+numbers[j] {
			return []int{i, j}
		}
		if target > numbers[i]+numbers[j] {
			i++
		}
		if target < numbers[i]+numbers[j] {
			j--
		}
	}
	return nil
}
