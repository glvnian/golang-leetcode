package day_14

import (
	"fmt"
	"testing"
)

func TestMinFlipsMonoIncr(t *testing.T) {
	s := "00110"
	s = "010110"
	r := minFlipsMonoIncr(s)
	fmt.Println(r)
}
