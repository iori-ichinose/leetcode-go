/*
 * https://leetcode-cn.com/problems/lemonade-change/
 * 2020.12.10
 * Golang 12ms(98.65%), 5.9MB(76.68%)
 */
package greedy

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, money := range bills {
		if money == 5 {
			five++
		} else if money == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
