package main

import (
	"fmt"
	"sort"
)

func main() {
	array1 := []int{0, 0, 0}
	array2 := []int{1, 0, -1}
	array3 := []int{-1, 0, 1, 2, -1, -4}
	array4 := []int{0, 1, 1}
	array5 := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("\n [0,0,0] output - %+v\n", threeSum(array1))
	fmt.Printf("\n [1, 0, -1] output - %+v\n", threeSum(array2))
	fmt.Printf("\n [-1, 0, 1, 2, -1, -4] output - %+v\n", threeSum(array3))
	fmt.Printf("\n [0,1,1] output - %+v\n", threeSum(array4))
	fmt.Printf("\n [-1,0,1,2,-1,-4] output - %+v\n", threeSum(array5))

}

func threeSum(nums []int) [][]int {
	var ret [][]int
	nlen := len(nums)
	sort.Ints(nums)
	for i := 0; i < nlen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := nlen - 1
		for start < end {
			if nums[start]+nums[end]+nums[i] == 0 {
				ele := []int{nums[start], nums[end], nums[i]}
				ret = append(ret, ele)
				start++
				end--
				for nums[start-1] == nums[start] && start < end {
					start++
				}
				for nums[end+1] == nums[end] && start < end {
					end--
				}
			}
			if nums[start]+nums[end]+nums[i] < 0 {
				start++
				for nums[start-1] == nums[start] && start < end {
					start++
				}
			}
			if nums[start]+nums[end]+nums[i] > 0 {
				end--
				for nums[end+1] == nums[end] && start < end {
					end--
				}
			}
		}
	}
	return ret
}

// func threeSum(nums []int) [][]int {
// 	validSum := 0
// 	triplets := make([][]int, 0)
// 	numsLength := len(nums)
// 	if numsLength < 3 || numsLength > 3000 {
// 		return triplets
// 	}
// 	for i := 0; i < numsLength; i++ {
// 		for j := i + 1; j < numsLength; j++ {
// 			for k := i + 2; k < numsLength; k++ {
// 				if i != j && j != k && k != i {
// 					if nums[i]+nums[j]+nums[k] == validSum {
// 						if nums[i] != nums[j] && nums[j] != nums[k] && nums[k] != nums[i] {
// 							arry := []int{nums[i], nums[j], nums[k]}
// 							triplets = append(triplets, arry)
// 						}
// 					}

// 				}
// 			}
// 		}
// 	}
// 	return triplets
// }
