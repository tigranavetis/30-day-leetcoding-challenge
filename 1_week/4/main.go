package main

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
		innerloop:
			for j := i; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break innerloop
				}
			}
		}
	}
}

func main() {
	nums := []int{10, 0, 0, 3, 12}
	moveZeroes(nums)
}
