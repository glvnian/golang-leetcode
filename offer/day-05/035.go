package day_05

import (
	"fmt"
	"time"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	var minSize int
	var m [1440]bool
	var minMinute, maxMinute int
	minMinute = len(m) + 1
	minSize = len(m) + 1
	for _, ts := range timePoints {
		var minute = GetMinute(ts)
		minMinute = Min(minMinute, minute)
		maxMinute = Max(maxMinute, minute)
		if m[minute] {
			return 0
		}
		m[minute] = true
	}

	var last int
	for i, ok := range m {
		if ok && last == 0 {
			last = i
			continue
		}
		if ok {
			minSize = Min(minSize, i-last)
			last = i
		}
	}
	minSize = Min(minSize, minMinute+1440-maxMinute)
	return minSize
}

func GetMinute(ts string) int {
	ti, err := time.Parse("15:04", ts)
	if err != nil {
		fmt.Println(err.Error())
	}
	h := ti.Hour()
	m := ti.Minute()
	return h*60 + m
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
