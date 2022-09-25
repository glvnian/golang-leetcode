package day_12

import (
	"fmt"
	"sort"
)

func mergeV1(intervals [][]int) [][]int {
	var r = make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		fmt.Println(i)
		tmp := intervals[i]
		var j = i + 1
		for ; j < len(intervals); j++ {
			if tmp[1] >= intervals[j][0] {
				tmp[1] = Max(tmp[1], intervals[j][1])
			} else {
				break
			}
		}
		r = append(r, tmp)
		i = j - 1
	}

	return r
}

func merge(intervals [][]int) [][]int {
	var r = make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		tmp := intervals[i]
		var j = i + 1
		for ; j < len(intervals) && tmp[1] >= intervals[j][0]; j++ {
			tmp[1] = Max(tmp[1], intervals[j][1])
		}
		r = append(r, tmp)
		i = j - 1
	}
	return r
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
