package qstring

import "fmt"

func strStr(haystack string, needle string) int {
	L, n := len(needle), len(haystack)
	if L == 0 {
		return 0
	}

	for pn := 0; pn < n-L+1; {
		for pn < n-L+1 && haystack[pn] != needle[0] {
			pn++
		}
		curLen, pL := 0, 0
		for pL < L && pn < n && haystack[pn] == needle[pL] {
			pn++
			pL++
			curLen++
		}

		if curLen == L {
			return pn - L
		}
		pn -= curLen - 1
	}

	return -1
}

func StrStrTb() {
	fmt.Print("Testbench strStr: ")
	fmt.Println(strStr("hello", "ll"))
}
