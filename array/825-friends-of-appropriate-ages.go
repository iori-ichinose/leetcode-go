/*
 * https://leetcode-cn.com/problems/friends-of-appropriate-ages/
 * 2020.11.22
 * Golang 36ms(90.48%), 6.7MB(23.81%)
 */
package array

func numFriendRequests(ages []int) int {
	var ageCount [121]int
	res := 0
	for _, n := range ages {
		ageCount[n]++
	}
	for i := 15; i <= 120; i++ {
		countA := ageCount[i]
		if countA == 0 {
			continue
		}
		for j := 15; j <= 120; j++ {
			countB := ageCount[j]
			if countB == 0 {
				continue
			}
			if int(float64(i)*0.5+7) >= j {
				continue
			} else if i < j {
				continue
			}
			if i == j {
				res += (countA - 1) * countB
			} else {
				res += countA * countB
			}
		}
	}
	return res
}

func NumFriendRequestsTb() {
	friends := []int{16, 15, 27, 120}
	println(numFriendRequests(friends))
}
