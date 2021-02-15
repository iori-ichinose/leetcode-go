package qstring

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	sign, abs := 1, s

	if s[0] == '+' {
		abs = s[1:]
	} else if s[0] == '-' {
		sign, abs = -1, s[1:]
	} else if !unicode.IsDigit(rune(s[0])) {
		return 0
	}

	for i, b := range abs {
		if !unicode.IsDigit(rune(b)) {
			abs = abs[:i]
			break
		}
	}

	res := 0
	for _, b := range abs {
		res = 10*res + int(b-'0')
		switch {
		case sign == 1 && res > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && res*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * res
}

func MyAtoiTb() {
	fmt.Print("TestBench myAtoi: ")
	fmt.Println(myAtoi("   -42"))
}
