package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("\nThe sum of 11 & 1 is  %v\n", addBinary("11", "1"))
	// fmt.Printf("\nThe sum of 1010 & 1011 is  %v\n", addBinary("1010", "1011"))
	// fmt.Printf("\nThe sum of 1010 & 1011 is  %v\n", addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
	// "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	res := ""
	aChars := []byte(a)
	bChars := []byte(b)
	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			aChar, _ := strconv.Atoi(string(aChars[i]))
			sum += aChar
			fmt.Println("sum in i", sum)
			i--
		}
		if j >= 0 {
			bChar, _ := strconv.Atoi(string(bChars[j]))
			sum += bChar
			fmt.Println("sum in j", sum)
			j--
		}
		fmt.Printf("Sum : %v", sum)
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
		res = fmt.Sprintf("%s%d", res, sum%2)
	}
	if carry != 0 {
		res = fmt.Sprintf("%s%d", res, carry)
	}
	fmt.Printf("res :%v", res)
	return reverse(res)
}
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
