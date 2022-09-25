package day_11

import (
	"fmt"
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	var piles []int
	var h int
	var r int
	//piles = []int{3, 6, 7, 11}
	//h = 8

	//piles = []int{30, 11, 23, 4, 20}
	piles = []int{30, 11, 23, 4, 20}
	h = 6

	//piles = []int{312884470}
	//h = 312884469

	r = minEatingSpeed(piles, h)
	fmt.Println(r)
}
