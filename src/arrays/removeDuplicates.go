package main

import "fmt"

func main() {
	array1 := []int{1, 1, 2}
	array2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	array3 := []int{1, 2, 2}
	fmt.Printf("\n array 1 {1, 1, 2} output : %d \n", removeDuplicates(array1))
	fmt.Printf("\n array 2 {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}output : %d \n", removeDuplicates(array2))
	fmt.Printf("\n array 2 [1,2,2]output : %d \n", removeDuplicates(array3))

}

func removeDuplicates(nums []int) int {
	numsLength := len(nums)
	numsMp := make(map[int]int, numsLength)
	if numsLength < 1 || numsLength > 30000 {
		return 0
	}
	for i := 0; i <= numsLength-1; i++ {
		for j := i; j <= numsLength-1; j++ {
			if nums[i] == nums[j] {
				continue
			} else {
				numsMp[nums[j]] = 1
				nums[i+1] = nums[j]
				i = i + 1
			}
		}
	}
	return len(numsMp) + 1
}

// func removeDuplicates(nums []int) int {
// 	expectedLength := 0
// 	numsLength := len(nums)
// 	numsMp := make(map[int]int, 0)
// 	if numsLength < 1 || numsLength > 30000 {
// 		return expectedLength
// 	}
// 	for i := 0; i < numsLength; i++ {
// 		//if nums[i] >= -100 || nums[i] <= 100 {
// 		numsMp[nums[i]] = 1
// 		//		}
// 	}
// 	j := 0
// 	for k, _ := range numsMp {
// 		fmt.Printf("\n Putting %d position vaalue : %d ", j, k)
// 		nums[j] = k
// 		j++
// 	}
// 	fmt.Println(numsMp)
// 	fmt.Println(nums)
// 	return len(numsMp)
// }

/// using two pointers we would need to solve this
