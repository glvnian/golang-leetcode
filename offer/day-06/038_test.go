package day_06

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	var temperatures, r []int
	temperatures = []int{73, 74, 75, 71, 69, 72, 76, 73}
	temperatures = []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	fmt.Println(temperatures)
	r = dailyTemperatures(temperatures)
	fmt.Println(r)
}

func TestT(t *testing.T) {
	var temperatures = []int{73, 74, 75, 71, 69, 72, 76, 73}
	a := make([]int, len(temperatures))
	fmt.Println(len(a))
}
