package _structures

type Stack interface {
	Push(elem interface{})
	Pop() interface{}
	Len() int
	IsEmpty() bool
	Top() interface{}
}

type ArrayStack struct {
	Stack
	data   []interface{}
	length int
}

func (this *ArrayStack) Push(elem interface{}) {
	this.data = append(this.data, elem)
	this.length++
}

func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	this.length--
	res := this.data[this.length]
	this.data = this.data[:this.length]
	return res
}

func (this *ArrayStack) Len() int {
	return this.length
}

func (this *ArrayStack) IsEmpty() bool {
	return this.length == 0
}

func (this *ArrayStack) Top() interface{} {
	return this.data[this.length-1]
}

type ListStack struct {
	Stack
	data List
}

func (this *ListStack) Push(elem interface{}) {
	this.data.AddBack(elem)
}

func (this *ListStack) Pop() interface{} {
	return this.data.RemoveBack()
}

func (this *ListStack) Len() int {
	return this.data.length
}

func (this *ListStack) IsEmpty() bool {
	return this.data.IsEmpty()
}

func (this *ListStack) Top() interface{} {
	return this.data.GetLastElem()
}
