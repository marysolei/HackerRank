package main

import "fmt"

func main() {
	n := 4
	staircase(n)
}

//Question Link: https://www.hackerrank.com/challenges/staircase/problem

func staircase(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}

}
