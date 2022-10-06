package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0})
	moveZeroes([]int{0, 0, 0, 1, 2, 3, 0, 0, 4, 5, 0, 10, 11})
	moveZeroes([]int{2, 1})
	moveZeroes([]int{1, 0, 1})
}

func moveZeroes(nums []int) {
	p := 0
	if len(nums) == 0 {
		return
	}
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}
}
