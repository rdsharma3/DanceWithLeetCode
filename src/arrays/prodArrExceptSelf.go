package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4}
	fmt.Println("product of [1,2,3,4]", productExceptSelf(array1))
}

//Brute Force with O(n^2)
// func productExceptSelf(nums []int) []int {

// 	if len(nums) < 2 || len(nums) > 100000 {
// 		return []int{}
// 	}
// 	prodNums := make([]int, len(nums))
// 	prod := 1
// 	for j := 0; j < len(nums); j++ {
// 		prod = 1
// 		for i := len(nums) - 1; i >= 0; i-- {
// 			if i == j {
// 				continue
// 			}
// 			prod *= nums[i]
// 		}
// 		prodNums[j] = prod
// 	}

// 	return prodNums
// }

func productExceptSelf(nums []int) []int {
	lengthNums := len(nums)
	leftPrefix := make([]int, lengthNums)
	rightPrefix := make([]int, lengthNums)
	ans := make([]int, lengthNums)
	for i := 0; i < lengthNums-1; i++ {
		if i == 0 {
			leftPrefix[i] = nums[0]
		} else {
			leftPrefix[i] = leftPrefix[i-1] * nums[i]
		}
	}
	for i := lengthNums - 1; i >= 0; i-- {
		if i == lengthNums-1 {
			rightPrefix[i] = nums[lengthNums-1]
		} else {
			rightPrefix[i] = rightPrefix[i+1] * nums[i]
		}
	}
	for i := 0; i < lengthNums; i++ {
		if i == 0 {
			ans[i] = rightPrefix[i+1]
		} else if i == lengthNums-1 {
			ans[i] = leftPrefix[i-1]
		} else {
			ans[i] = leftPrefix[i-1] * rightPrefix[i+1]
		}
	}
	return ans
}

//Optimised sol
// public int[] productExceptSelf(int[] nums) {
// 	int[] leftPrefix = new int[nums.length];
// 	int[] rightPrefix = new int[nums.length];
// 	int[] ans = new int[nums.length];

// 	for(int i = 0;i<nums.length;i++){
// 		if(i==0) leftPrefix[i] = nums[0];
// 		else{
// 			leftPrefix[i] = leftPrefix[i-1]*nums[i];
// 		}
// 	}

// 	 for(int i = nums.length-1;i>=0;i--){
// 		if(i==nums.length-1) rightPrefix[i] = nums[nums.length-1];
// 		else{
// 			rightPrefix[i] = rightPrefix[i+1]*nums[i];
// 		}
// 	}
// 	 for(int i = 0;i<nums.length;i++){
// 		 if(i==0){
// 			 ans[i] = rightPrefix[i+1];
// 		 }else if(i==nums.length-1){
// 			 ans[i] = leftPrefix[i-1];
// 		 }else{
// 		 ans[i] = leftPrefix[i-1] * rightPrefix[i+1];
// 		 }
// 	 }
// 	 return ans;
//  }
