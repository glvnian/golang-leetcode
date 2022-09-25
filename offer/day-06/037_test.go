package day_06

import (
	"fmt"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	var asteroids, r []int
	//asteroids = []int{5, 10, -5}
	//asteroids = []int{5, -5}
	//asteroids = []int{10, 2, -5}
	asteroids = []int{-2, -2, 1, -2}

	r = asteroidCollision(asteroids)
	fmt.Println(r)
}

func TestTest(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(a[:3-1])
}
