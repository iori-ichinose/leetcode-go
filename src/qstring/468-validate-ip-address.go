/*
 * https://leetcode-cn.com/problems/validate-ip-address/
 * 2020.11.22
 * Golang 0ms(100.00%), 1.9MB(95.52%)
 */
package qstring

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	s := strings.Split(IP, ".")
	if len(s) == 4 {
		IPv4 := func(nums []string) string {
			isDigit := func(str string) bool {
				for _, n := range str {
					if n < '0' || n > '9' {
						return false
					}
				}
				return true
			}
			atoi := func(str string) int {
				n, _ := strconv.Atoi(str)
				return n
			}
			for i := range nums {
				if len(nums[i]) == 0 || len(nums[i]) > 3 {
					return "Neither"
				}
				if nums[i][0] == '0' && len(nums[i]) != 1 ||
					!isDigit(nums[i]) || atoi(nums[i]) > 255 {
					return "Neither"
				}
			}
			return "IPv4"
		}
		return IPv4(s)
	} else {
		IPv6 := func(nums []string) string {
			if len(nums) != 8 {
				return "Neither"
			}
			valid := "0123456789ABCDEFabcdef"
			include := func(ch rune) bool {
				for _, n := range valid {
					if n == ch {
						return true
					}
				}
				return false
			}
			for i := range nums {
				if len(nums[i]) == 0 || len(nums[i]) > 4 {
					return "Neither"
				}
				for _, c := range nums[i] {
					if !include(c) {
						return "Neither"
					}
				}
			}
			return "IPv6"
		}
		return IPv6(strings.Split(IP, ":"))
	}
}

func ValidIPAddressTb() {
	fmt.Println(validIPAddress("172.16.254.8"))
}
