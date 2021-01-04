package slideWindow

type Deque struct {
	data []int
	size int
}

func (d *Deque) pushBack(elem int) {
	d.data = append(d.data, elem)
	d.size++
}

func (d *Deque) getBack() int {
	return d.data[d.size-1]
}

func (d *Deque) getFront() int {
	return d.data[0]
}

func (d *Deque) popBack() int {
	d.size--
	res := d.data[d.size]
	d.data = d.data[:d.size]
	return res
}

func (d *Deque) popFront() int {
	d.size--
	res := d.data[0]
	d.data = d.data[1:]
	return res
}

func (d *Deque) isEmpty() bool {
	return d.size == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	var window Deque
	res := make([]int, n-k+1)

	for i := 0; i < n; i++ {
		if !window.isEmpty() && window.getFront() <= i-k {
			window.popFront()
		}
		for !window.isEmpty() && nums[i] > nums[window.getBack()] {
			window.popBack()
		}
		window.pushBack(i)
		if i >= k-1 {
			res[i-k+1] = nums[window.getFront()]
		}
	}
	return res
}
