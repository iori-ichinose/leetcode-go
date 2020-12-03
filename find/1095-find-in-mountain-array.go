/*
 * https://leetcode-cn.com/problems/find-in-mountain-array/
 * 2020.12.2
 * Golang 4ms(92.86%), 3.3MB(78.57%)
 */
package find

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 0 }

func findMountainTop(array *MountainArray, left, right int) int {
	for left < right {
		mid := left + (right-left)/2
		if array.get(mid) < array.get(mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func findInArray(array *MountainArray, left, right int, target int) int {
	for left < right {
		mid := left + (right-left)/2
		if array.get(mid) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if array.get(left) == target {
		return left
	} else {
		return -1
	}
}

func findInReversedArray(array *MountainArray, left, right int, target int) int {
	for left < right {
		mid := left + (right-left+1)/2
		if array.get(mid) < target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	if array.get(left) == target {
		return left
	} else {
		return -1
	}
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	top := findMountainTop(mountainArr, 0, n-1)
	if mountainArr.get(top) == target {
		return top
	}

	findLeft := findInArray(mountainArr, 0, top-1, target)
	if findLeft != -1 {
		return findLeft
	} else {
		return findInReversedArray(mountainArr, top+1, n-1, target)
	}
}
