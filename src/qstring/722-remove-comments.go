/*
 * https://leetcode-cn.com/problems/remove-comments/
 * 2020.11.22
 * Golang 0ms(100.00%), 2.1MB(28.57%)
 */
package qstring

import "fmt"

func removeComments(source []string) []string {
	inBlock := false
	var res []string
	curLine := ""
	for _, line := range source {
		i := 0
		if !inBlock {
			curLine = ""
		}
		for i < len(line) {
			if !inBlock && i < len(line)-1 && line[i] == '/' && line[i+1] == '*' {
				inBlock = true
				i += 2
			} else if inBlock && i < len(line)-1 && line[i] == '*' && line[i+1] == '/' {
				inBlock = false
				i += 2
			} else if !inBlock && i < len(line)-1 && line[i] == '/' && line[i+1] == '/' {
				break
			} else if !inBlock {
				curLine += string(line[i])
				i++
			} else {
				i++
			}
		}

		if curLine != "" && !inBlock {
			res = append(res, curLine)
		}
	}
	return res
}

func RemoveCommentsTb() {
	fmt.Println("Testbench removeComments: ")
	strs := []string{
		"/*Test program */",
		"int main()",
		"{ ",
		"  // variable declaration ",
		"int a, b, c;",
		"/* This is a test",
		"   multiline  ",
		"   comment for ",
		"   testing */",
		"a = b + c;", "}",
	}
	for _, n := range removeComments(strs) {
		fmt.Println(n)
	}
}
