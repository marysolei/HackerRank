package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/compare-the-triplets/problem

func main() {
	a := []int32{17, 28, 39}
	b := []int32{99, 16, 8}
	fmt.Println(compareTriplets(a, b))
}

func compareTriplets(a []int32, b []int32) []int32 {

	alicepnt, bobpnt := int32(0), int32(0)
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alicepnt++
		} else if a[i] < b[i] {
			bobpnt++
		}
	}
	return []int32{alicepnt, bobpnt}

}
