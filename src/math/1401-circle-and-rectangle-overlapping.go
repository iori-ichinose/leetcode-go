/*
 * https://leetcode-cn.com/problems/circle-and-rectangle-overlapping/
 * 2020.11.26
 * Golang 0ms(100.00%), 2MB(60.00%)
 */
package math

import "math"

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	if x_center < x1-radius || x_center > x2+radius ||
		y_center < y1-radius || y_center > y2+radius {
		return false
	}

	centerDis := math.Sqrt(float64((x_center-(x2+x1)/2)*(x_center-(x2+x1)/2) + (y_center-(y2+y1)/2)*(y_center-(y2+y1)/2)))
	radDis := math.Sqrt(float64((x2-x1)*(x2-x1)+(y2-y1)*(y2-y1)))/2 + float64(radius)
	return centerDis <= radDis
}
