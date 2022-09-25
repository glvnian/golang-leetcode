package day_01

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	str1 := "101010"
	str2 := "101010"
	a := addBinary(str1, str2)
	fmt.Println(a)
}
