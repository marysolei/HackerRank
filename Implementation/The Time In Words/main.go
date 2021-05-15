package main

import "fmt"

func main() {

	fmt.Println(timeInWords(11, 40))
}
func timeInWords(h int32, m int32) string {
	h = h % 12
	hourCount := []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
		"twenty",
		"twenty one",
		"twenty two",
		"twenty three",
		"twenty four",
		"twenty five",
		"twenty six",
		"twenty seven",
		"twenty eight",
		"twenty nine"}

	if m == 0 {
		return hourCount[h] + " o' clock"

	} else if m == 15 {

		return "quarter past " + hourCount[h]

	} else if m == 30 {
		return "half past " + hourCount[h]
	} else if m == 45 {
		return "quarter to " + hourCount[h+1]
	} else if m < 30 {

		return hourCount[m] + " minutes past " + hourCount[h]

	}

	return hourCount[60-m] + " minutes to " + hourCount[h+1]
}
