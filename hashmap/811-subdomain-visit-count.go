/*
 * https://leetcode-cn.com/problems/subdomain-visit-count/
 * 2020.12.2
 * Golang 12ms(55.36%), 6.4MB(21.43%)
 */
package hashmap

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var res []string
	count := map[string]int{}
	for i := range cpdomains {
		dom := strings.Split(cpdomains[i], " ")
		num, _ := strconv.Atoi(dom[0])
		urls := strings.Split(dom[1], ".")
		for j := range urls {
			count[strings.Join(urls[j:], ".")] += num
		}
	}

	for k, v := range count {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}
