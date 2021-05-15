//Question's Link:https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println(beautifulDays(20, 23, 6))
}

func beautifulDays(i int32, j int32, k int32) int32 {
	var res string
	var cnt int32
	for e := i; e <= j; e++ {
		tmpstr := strconv.Itoa(int(e))
		for _, v := range tmpstr {
			res = string(v) + res
		}
		erev, _ := strconv.Atoi(res)
		res = ""
		if int32(math.Abs(float64(e)-float64(erev)))%(k) == 0 {
			cnt++
		}
	}
	return cnt
}
