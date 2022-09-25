package day_11

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	var w []int
	var r int
	w = []int{1, 3}
	obj := Constructor(w)
	r = obj.PickIndex()
	fmt.Println(r)

}

func TestName(t *testing.T) {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(3)
	fmt.Println(num)
}

func TestPick(t *testing.T) {
	var w = []int{1, 3}
	var r = pick(w, 1)
	fmt.Println(r)
}
