package day_14

import (
	"fmt"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	s1 := "rabbbit"
	s2 := "rabbit"
	r := numDistinct(s1, s2)
	fmt.Println(r)
}
