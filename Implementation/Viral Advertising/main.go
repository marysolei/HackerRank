package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(viralAdvertising(3))
}

//Question Link: https://www.hackerrank.com/challenges/strange-advertising/problem

func viralAdvertising(n int32) int32 {
	shared := 5
	cum := 0
	for i := 1; i <= int(n); i++ {
		liked := math.Floor(float64(shared / 2))
		cum += int(liked)
		shared = int(liked) * 3
	}
	return int32(cum)
}
