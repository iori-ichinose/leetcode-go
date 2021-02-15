package _structures

import "fmt"

type Deque interface {
	Empty() bool
	Size() int
	PushBack(elem interface{})
	PushFront(elem interface{})
	GetFront() interface{}
	GetBack() interface{}
	PopFront() interface{}
	PopBack() interface{}

	toString() string
}

type ArrayDeque struct {
	Deque
	data []interface{}
	size int
}

func (this *ArrayDeque) Empty() bool {
	return this.size == 0
}

func (this *ArrayDeque) Size() int {
	return this.size
}

func (this *ArrayDeque) PushBack(elem interface{}) {
	this.data = append(this.data, elem)
	this.size++
}

func (this *ArrayDeque) PushFront(elem interface{}) {
	this.data = append([]interface{}{elem}, this.data...)
	this.size++
}

func (this *ArrayDeque) GetFront() interface{} {
	return this.data[0]
}

func (this *ArrayDeque) GetBack() interface{} {
	return this.data[this.size-1]
}

func (this *ArrayDeque) PopFront() interface{} {
	this.size--
	res := this.data[0]
	this.data = this.data[1:]
	return res
}

func (this *ArrayDeque) PopBack() interface{} {
	this.size--
	res := this.data[this.size]
	this.data = this.data[:this.size]
	return res
}

func (this *ArrayDeque) toString() string {
	return fmt.Sprintf("ArrayDeque at %p", this)
}

func __dequeTest() {
	var s Deque = new(ArrayDeque)
	for i := 0; i < 10; i++ {
		s.PushBack(i)
	}
	for !s.Empty() {
		print(s.PopBack().(int))
	}
	for i := 0; i < 10; i++ {
		s.PushBack(i)
	}
	for !s.Empty() {
		print(s.PopFront().(int))
	}
	println(s.toString())
}
