package day_14

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	r := coinChange(coins, amount)
	fmt.Println(r)

}
