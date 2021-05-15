package main

//Question Link: https://www.hackerrank.com/challenges/cats-and-a-mouse/problems

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(catAndMouse(2, 5, 4))

}

func catAndMouse(x int32, y int32, z int32) string {

	diffA := math.Abs(float64(x) - float64(z))
	diffB := math.Abs(float64(y) - float64(z))
	if diffA < diffB {
		return "Cat A"
	} else if diffB < diffA {
		return "Cat B"
	}
	return "Mouse C"
}
