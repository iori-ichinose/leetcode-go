package hashmap

func powerfulIntegers(x int, y int, bound int) []int {
	res := map[int]int{}
	var helper func(int, int, int, int)
	helper = func(i, j, numx, numy int) {
		num := numx + numy
		if num > bound {
			return
		}
		res[num] = 1
		if y != 1 {
			helper(i, j+1, numx, y*numy)
		}

		if x != 1 {
			helper(i+1, j, x*numx, numy)
		}
	}

	helper(0, 0, 1, 1)
	var ans []int
	for i := range res {
		ans = append(ans, i)
	}
	return ans
}

func PowerfulIntegersTb() []int {
	return powerfulIntegers(3, 5, 15)
}
