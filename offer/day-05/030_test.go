package day_05

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestConstructor(t *testing.T) {
	var val, res int
	var b bool
	obj := Constructor()

	val = 1
	b = obj.Insert(val)
	fmt.Println("b1: ", b)
	val = 2
	b = obj.Insert(val)
	fmt.Println("b2: ", b)
	b = obj.Remove(val)
	fmt.Println("b2: ", b)
	res = obj.GetRandom()
	fmt.Println("res: ", res)

}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(1)
	fmt.Println(a)
	//a := []int{1, 2, 3}
	//fmt.Println(a)
	//a = a[:2]
	//fmt.Println(a)
}
