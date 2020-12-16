/*
 * https://leetcode-cn.com/problems/monotone-increasing-digits/
 * 2020.12.15
 * Golang 0ms(100.00%), 2MB(80.00%)
 */
package hashmap

func maxNumberOfBalloons(text string) int {
	count := [5]int{}
	for _, ch := range text {
		switch ch {
		case 'b':
			count[0]++
			break
		case 'a':
			count[1]++
			break
		case 'l':
			count[2]++
			break
		case 'o':
			count[3]++
			break
		case 'n':
			count[4]++
			break
		}
	}

	res := 1<<31 - 1
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	for _, ch := range "balon" {
		if ch == 'l' || ch == 'o' {
			res = min(res, count[ch]/2)
		} else {
			res = min(res, count[ch])
		}
	}
	return res
}
