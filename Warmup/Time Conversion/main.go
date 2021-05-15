package main

import (
	"fmt"
	"strconv"
)

//Question Link: https://www.hackerrank.com/challenges/time-conversion/problem

func main() {
	s := "12:05:45AM"
	fmt.Println(timeConversion(s))

}

func timeConversion(s string) string {
	if s[8:] == "AM" && s[:2] == "12" {
		return "00" + s[2:8]
	} else if s[8:] == "PM" && s[:2] == "12" {
		return s[:8]
	} else if s[8:] == "PM" {
		tmp, _ := strconv.Atoi(s[:2])
		tmp += 12
		return strconv.Itoa(tmp) + s[2:8]
	}
	return s[:8]
}
