package day_12

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFindKthLargest(t *testing.T) {
	var arr1 []int
	var k, r int
	arr1 = []int{3, 2, 1, 5, 6, 4}
	k = 2
	r = findKthLargest(arr1, k)
	fmt.Println(r)

}

func TestTest(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(2)
	fmt.Println(a)
}
