package day_10

import (
	"fmt"
	"testing"
)

func TestMinimumLengthEncoding(t *testing.T) {
	var words = []string{"time", "me", "bell"}
	r := minimumLengthEncoding(words)
	fmt.Println(r)
}
