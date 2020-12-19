/*
 * https://leetcode-cn.com/problems/unique-email-addresses/
 * 2020.12.17
 * Golang 8ms(89.36%), 5.4MB(87.23%)
 */
package qstring

func numUniqueEmails(emails []string) int {
	mails := map[string]bool{}
	for i := range emails {
		var res []byte
		j := 0
		for ; emails[i][j] != '@' && emails[i][j] != '+'; j++ {
			if emails[i][j] != '.' {
				res = append(res, emails[i][j])
			}
		}
		for emails[i][j] != '@' {
			j++
		}
		mails[string(res)+emails[i][j:]] = true
	}
	return len(mails)
}
