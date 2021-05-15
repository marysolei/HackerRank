package main

import (
	"fmt"
)

//Question Link: https://www.hackerrank.com/challenges/migratory-birds/problem

func main() {
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	fmt.Println(migratoryBirds(arr))
}

func migratoryBirds(arr []int32) int32 {
	birdType := make([]int32, 5)
	var index int
	var max int32
	for _, e := range arr {
		birdType[e-1]++
	}
	for i := 0; i < 5; i++ {
		if birdType[i] > max {
			max = birdType[i]
			index = i
		}
	}
	return int32(index + 1)

}
