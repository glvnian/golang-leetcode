package day_05

import (
	"fmt"
	"testing"
)

func TestFindMinDifference(t *testing.T) {
	var timePoints []string
	var r int
	timePoints = []string{"00:00", "23:59", "00:00"}
	fmt.Println(timePoints)
	r = findMinDifference(timePoints)
	fmt.Println(r)
}

func TestGetMinute(t *testing.T) {

}
