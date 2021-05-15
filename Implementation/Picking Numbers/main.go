package main

import (
	"fmt"
	"sort"
)

// Question Link: https://www.hackerrank.com/challenges/picking-numbers/problem

func main() {

	p := []int32{4, 6, 5, 3, 3, 1}
	fmt.Println(pickingNumbers(p))
}

func pickingNumbers(a []int32) int32 {

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	l := len(a)
	var sameEl int32
	maxv := 0
	for i := l - 1; i >= 0; i-- {
		cnt := 1
		if sameEl == a[i] {
			continue
		} else {
			sameEl = a[i]
		}
		for j := i - 1; j >= 0; j-- {
			if a[i]-a[j] < 2 {
				cnt++
			}
		}
		if cnt > maxv {
			maxv = cnt
		}
	}
	return int32(maxv)

}
