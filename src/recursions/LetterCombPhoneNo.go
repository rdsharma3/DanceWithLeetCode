package main

import "fmt"

func main() {
	fmt.Println("2,3:", letterCombinations("23"))
	fmt.Println("7,8,9:", letterCombinations("789"))
}

var digitLetterDict = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

///****** using iteration ****
// func letterCombinations(digits string) []string {
// 	var result []string
// 	if digits == "" {
// 		return result
// 	}
// 	result = []string{""}
// 	for i := range digits {
// 		sts := digitLetterDict[digits[i]]
// 		fmt.Println("sts:", sts)
// 		var next []string
// 		for v := range sts {
// 			c := sts[v]
// 			for _, r := range result {
// 				next = append(next, r+string(c))
// 			}
// 		}
// 		result = next
// 	}
// 	return result
// }

/// **** Using Recursion as below ******

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	recurHelper(digits, 0, &result, "")
	return result
}

func recurHelper(digits string, p int, result *[]string, s string) {
	if p == len(digits) {
		fmt.Println("p : ", p)
		fmt.Println("digits : ", digits)
		*result = append(*result, s)
		return
	}
	strs := digitLetterDict[digits[p]]
	for i := range strs {
		fmt.Println("strs : ", strs)
		fmt.Println("s+string(strs[i]) : ", s+string(strs[i]))
		recurHelper(digits, p+1, result, s+string(strs[i]))
	}
}
