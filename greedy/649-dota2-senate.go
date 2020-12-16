/*
 * https://leetcode-cn.com/problems/dota2-senate/
 * 2020.12.11
 * Golang 4ms(60.00%), 2.4MB(100.00%)
 */
package greedy

func predictPartyVictory(senate string) string {
	if len(senate) == 0 {
		return ""
	}
	seList := []byte(senate)
	R, D, person := true, true, 0
	for R && D {
		R, D = false, false
		for i := range seList {
			if seList[i] == 'R' {
				R = true
				if person < 0 {
					seList[i] = '0'
				}
				person++
			} else if seList[i] == 'D' {
				D = true
				if person > 0 {
					seList[i] = '0'
				}
				person--
			}
		}
	}
	if person > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func PredictPartyVictoryTb() {
	print("Testbench predictPartyVictory: ")
	println(predictPartyVictory("RD"))
}
