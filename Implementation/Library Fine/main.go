package main

import (
	"fmt"
)

//Question Link: https://www.hackerrank.com/challenges/library-fine/problem

func main() {
	fmt.Println(libraryFine(1, 12, 2017, 1, 12, 2017))
}

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {

	var fine int32
	if y2 < y1 {
		fine = 10000
	} else if y1 == y2 {
		if m2 < m1 {
			fine = 500 * (m1 - m2)
		} else if m2 == m1 {
			if d2 < d1 {
				fine = 15 * (d1 - d2)
			}
		}
	}
	return fine
}
