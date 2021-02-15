package array

import "fmt"

func moveZeroes(nums []int) {
	i := 0
	for _, n := range nums {
		if n != 0 {
			nums[i] = n
			i++
		}
	}
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}

func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func MoveZeroesTb() {
	fmt.Print("Testbench moveZeroesTb: ")
	test := []int{0, 1, 0, 3, 12}
	moveZeroes(test)
	for _, n := range test {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
