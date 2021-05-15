package main

import (
	"fmt"
	"math"
)

//Question Link: https://www.hackerrank.com/challenges/sherlock-and-squares/problem

func main() {

	fmt.Println(squares(3, 9))
}

func squares(a int32, b int32) int32 {
	var cnt int32
	low := int(math.Sqrt(float64(a)))

	for i := low; math.Pow(float64(i), 2) <= float64(b); i++ {
		if math.Pow(float64(i), 2) >= float64(a) {
			cnt++
		}
	}
	return cnt
}
