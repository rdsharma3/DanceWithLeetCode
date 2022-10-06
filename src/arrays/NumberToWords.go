package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("123 :", numberToWords(123))
	fmt.Println("12345 :", numberToWords(12345))
	fmt.Println("1234567 :", numberToWords(1234567))
	fmt.Println("50868 :", numberToWords(50868))
	fmt.Println("30888 :", numberToWords(30888))
}

var (
	Tens = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	Ones = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven",
		"Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
)

func numberToWords(num int) string {
	output := "Zero"
	if num == 0 {
		return output
	}
	return strings.Trim(recurExtractor(num), " ")
}

func recurExtractor(num int) string {
	if num >= 1000000000 {
		return strings.Trim(recurExtractor(num/1000000000)+" Billion "+recurExtractor(num%1000000000), " ")
	}
	if num >= 1000000 {
		return strings.Trim(recurExtractor(num/1000000)+" Million "+recurExtractor(num%1000000), " ")
	}
	if num >= 1000 {
		return strings.Trim(recurExtractor(num/1000)+" Thousand "+recurExtractor(num%1000), " ")
	}
	if num >= 100 {
		return strings.Trim(recurExtractor(num/100)+" Hundred "+recurExtractor((num%100)), " ")
	}
	if num >= 20 {
		return strings.Trim(Tens[num/10]+" "+recurExtractor(num%10), " ")
	}
	return Ones[num]
}
