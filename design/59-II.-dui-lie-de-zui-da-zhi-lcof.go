package design

type MaxQueue struct {
	data    []int
	maxData []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.maxData[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	for len(this.maxData) != 0 && value > this.maxData[len(this.maxData)-1] {
		this.maxData = this.maxData[:len(this.maxData)-1]
	}
	this.maxData = append(this.maxData, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	n := this.data[0]
	this.data = this.data[1:]
	if this.maxData[0] == n {
		this.maxData = this.maxData[1:]
	}
	return n
}
