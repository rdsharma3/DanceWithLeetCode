package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//myAtoi("432")
	// ret := myAtoi(" -432 ")
	// fmt.Println("Actual Return", ret)
	// ret1 := myAtoi("4193 with words")
	// fmt.Println("Actual Return", ret1)
	// ret1 = myAtoi("+1")
	// fmt.Println("Actual Return", ret1)
	// ret1 = myAtoi("+-12")
	// fmt.Println("Actual Return", ret1)
	// ret1 = myAtoi("00000-42a1234")
	// fmt.Println("Actual Return", ret1)
	// ret1 = myAtoi("words 123123")
	// fmt.Println("Actual Return", ret1)
	// ret1 = myAtoi("3.1213123")
	// fmt.Println("Actual Return", ret1)
	fmt.Println(myAtoi("00000-42a1234"))
	fmt.Println(myAtoi("  -0012a42"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("-9223372036854775809"))
}

func myAtoi(s string) int {
	retNum := 0
	nums := []byte{}

	s = strings.TrimLeft(s, " ")

	chars := []byte(s)

	isNegative := false
	isPositive := false
	for index, char := range chars {

		if string(char) == "-" {
			if isPositive || index > 0 {
				break
			}
			isNegative = true
		} else if string(char) == "+" {
			if isNegative || index > 0 {
				break
			}
			isPositive = true
		} else if char >= 48 && char <= 57 {
			nums = append(nums, char)
		} else {
			break
		}
	}
	for _, num := range nums {
		num -= '0'
		retNum = retNum*10 + int(num)
		if retNum > math.MaxInt32 && !isNegative {
			return math.MaxInt32
		}
		if retNum < math.MinInt32 && isNegative {
			return math.MinInt32
		}
	}
	if isNegative {
		retNum = -retNum
	}
	if retNum < math.MinInt32 {
		return math.MinInt32
	}
	return retNum
}
