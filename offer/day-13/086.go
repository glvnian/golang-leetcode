package day_13

func partition(s string) [][]string {
	result := make([][]string, 0)
	subnet := []string{}
	helperPartition(s, 0, subnet, &result)
	return result
}

func helperPartition(s string, i int, subnet []string, result *[][]string) {
	if i == len(s) {
		*result = append(*result, append([]string{}, subnet...))
	} else {
		for j := i; j < len(s); j++ {
			if isPartition(s, i, j) {
				subnet = append(subnet, s[i:j+1])
				helperPartition(s, j+1, subnet, result)
				subnet = subnet[:len(subnet)-1]
			}
		}
	}
}

func isPartition(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
