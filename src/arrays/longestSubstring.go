package main

import "fmt"

func main() {
	lengthOfLongestSubstring("abcdefa")
	lengthOfLongestSubstring("abcabcbb")
	lengthOfLongestSubstring("bbbbbb")
	lengthOfLongestSubstring("pwwkew")
	lengthOfLongestSubstring(" ")
}

func lengthOfLongestSubstring(str string) (string, int) {
	longestStr := ""
	length := 0
	substrings := getSubStrings(str)
	longestStr, length = findLongest(substrings)
	fmt.Printf("\nReturning string : '%s' with length : %v\n", longestStr, length)
	return longestStr, length
}

func getSubStrings(str string) map[string]int {
	substrings := map[string]int{}
	strArray := []byte(str)
	for i := 0; i <= len(strArray)-1; i++ {
		substring := ""
		for j := i; j <= len(strArray)-1; j++ {
			if !contains(substring, strArray[j]) {
				substring += string(strArray[j])
			} else {
				i = j - 1
				break
			}
		}
		substrings[substring] = len(substring)
	}
	return substrings
}

func findLongest(substrings map[string]int) (string, int) {
	longest := 0
	longestStr := ""
	for k, v := range substrings {
		if v > longest {
			longest = v
			longestStr = k
		} else {
			continue
		}
	}
	return longestStr, longest
}

func contains(substring string, char byte) bool {
	if len(substring) == 0 {
		return false
	}
	strArray := []byte(substring)
	for i := 0; i <= len(strArray)-1; i++ {
		if string(strArray[i]) == string(char) {
			return true
		} else {
			continue
		}
	}
	return false
}
