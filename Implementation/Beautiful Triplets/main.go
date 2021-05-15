package main

import (
	"fmt"
)

//Question Link: https://www.hackerrank.com/challenges/beautiful-triplets/problem

func main() {

	//a := []int32{1, 2, 4, 5, 7, 8, 10}
	a := []int32{2, 2, 3, 4, 5}
	fmt.Println(beautifulTriplets(1, a))
}

func beautifulTriplets(d int32, arr []int32) int32 {
	var cnt int32
	for _, e := range arr {
		if contains(arr, (e+d)) && contains(arr, (e+2*d)) {
			cnt++
		}
	}
	return cnt
}

func contains(ar []int32, e int32) bool {

	for _, v := range ar {
		if v == e {
			return true
		}
	}
	return false
}
