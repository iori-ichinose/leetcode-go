/*
 * https://leetcode-cn.com/problems/goal-parser-interpretation/
 * 2020.12.6
 * Golang 0ms(100.00%), 2MB(100.00%)
 */
package qstring

func interpret(command string) string {
	var res []byte
	length := len(command)
	for i := 0; i < length; i++ {
		if command[i] == 'G' {
			res = append(res, byte('G'))
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				res = append(res, byte('o'))
				i++
			} else {
				res = append(res, byte('a'), byte('l'))
				i += 3
			}
		}
	}
	return string(res)
}
