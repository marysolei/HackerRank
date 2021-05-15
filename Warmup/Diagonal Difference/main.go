package main

import "math"

//Question Link: https://www.hackerrank.com/challenges/diagonal-difference/problem

func main() {

}

//O(n^2)
func diagonalDifference(arr [][]int32) int32 {
	var sumd1 int32
	var sumd2 int32
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				sumd1 += arr[i][j]
			}
			if i+j == n-1 {
				sumd2 += arr[i][j]
			}
		}
	}
	return int32(math.Abs(float64(sumd1) - float64(sumd2)))

}

//O(n)
func diagonalDifferenceON(arr [][]int32) int32 {
	var sumd1 int32
	var sumd2 int32
	n := len(arr)
	for i := 0; i < n; i++ {
		sumd1 += arr[i][i]
		sumd2 += arr[i][n-i-1]
	}

	return int32(math.Abs(float64(sumd1) - float64(sumd2)))

}
