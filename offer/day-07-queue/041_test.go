package day_07_queue

import (
	"fmt"
	"testing"
)

func TestMovingAverage(t *testing.T) {
	var r float64
	movingAverage := Constructor(3)
	r = movingAverage.Next(1) // 返回 1.0 = 1 / 1
	fmt.Println(r)
	r = movingAverage.Next(10) // 返回 5.5 = (1 + 10) / 2
	fmt.Println(r)
	r = movingAverage.Next(3) // 返回 4.66667 = (1 + 10 + 3) / 3
	fmt.Println(r)
	r = movingAverage.Next(5) // 返回 6.0 = (10 + 3 + 5) / 3
	fmt.Println(r)
}
