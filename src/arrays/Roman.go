package main

import "fmt"

var RomanNumeralDict map[string]int

func main() {
	//RomanNumeralDict = romanNumeralDict()
	//str := "CVII"
	//str := "III"
	//romanToInt(str)
	fmt.Println(intToRoman(50))
	fmt.Println(intToRoman(49))
	fmt.Println(intToRoman(499))
}

var romanNumeralDict = func() map[string]int {
	return map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

var numeralRomanDict = func() map[int]string {
	return map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
}

func romanToInt(str string) {
	romanNumeralDict()
	sum := 0
	chars := []byte(str)
	lenStr := len(str)
	for i := 0; i < lenStr; {
		if i < lenStr-1 && RomanNumeralDict[string(chars[i+1])] > RomanNumeralDict[string(chars[i])] {
			fmt.Printf("\nroman RomanNumeralDict[chars[i]] : %d , RomanNumeralDict[chars[i+1]] : %d \n", RomanNumeralDict[string(chars[i])], RomanNumeralDict[string(chars[i+1])])
			sum += RomanNumeralDict[string(chars[i+1])] - RomanNumeralDict[string(chars[i])]
			i += 2
		} else {
			sum = sum + RomanNumeralDict[string(chars[i])]

			i += 1
		}
	}
	println("Output :", sum)
}

var intRomanDict map[int]string

func intToRoman(num int) string {
	var amount int
	output := ""
	divisors := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < 13; i++ {
		amount = num / divisors[i]
		for j := 0; j < amount; j++ {
			output += romans[i]
		}
		num = num - amount*divisors[i]
	}
	return output
}
