package main

import (
	"fmt"
	"math"
)

//Question Link: https://www.hackerrank.com/challenges/utopian-tree/problem

func main() {
	fmt.Println(utopianTree(4))
}

func utopianTree(n int32) int32 {
	var res float64
	if n%2 == 0 {
		res = math.Pow(2, (float64(n)/2))*2 - 1
	} else {
		res = (math.Pow(2, (float64(n)-1)/2)*2 - 1) * 2
	}
	return int32(res)

}
