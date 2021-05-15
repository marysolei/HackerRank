// Question Link: https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problems

package main

import "fmt"

func main() {

	//arr := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	arr := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	fmt.Println(breakingRecords(arr))

}

func breakingRecords(scores []int32) []int32 {
	maxtmp := scores[0]
	mintmp := scores[0]
	maxhit := 0
	minhit := 0
	for _, score := range scores {
		if score > maxtmp {
			maxtmp = score
			maxhit++
		} else if score < mintmp {
			mintmp = score
			minhit++
		}
	}
	return []int32{int32(maxhit), int32(minhit)}
}
