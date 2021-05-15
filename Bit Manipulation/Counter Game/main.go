package main

import (
	"fmt"
	"math"
)

//Question Link: https://www.hackerrank.com/challenges/counter-game/problem

func main() {
	n := int64(30)
	fmt.Println(counterGame(n))
}

func counterGame(n int64) string {
	cnt := 0

	newnum := n
	for newnum > 1 {
		rt := math.Log2(float64(newnum))
		if rt-math.Floor(rt) > 0 {
			base := math.Floor(rt)
			newnum = int64(float64(newnum) - math.Pow(2, base))
			cnt++
		} else {
			newnum /= 2
			cnt++
		}
	}
	if newnum == 1 && cnt%2 != 0 {
		return "Louise"
	}
	return "Richard"
}
