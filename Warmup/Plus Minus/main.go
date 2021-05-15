package main

import (
	"fmt"
	"math"
)

//Question Link: https://www.hackerrank.com/challenges/plus-minus/problem

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusminus(arr)
}

func plusminus(arr []int32) {
	var sum int32
	max := int32(0)
	min := int32(math.MaxInt32)
	for _, v := range arr {
		sum += v
	}
	for i, _ := range arr {
		s := sum - arr[i]
		if max < s {
			max = s
		}
		if min > s {
			min = s
		}

	}
	fmt.Println(sum-max, sum-min)
}
