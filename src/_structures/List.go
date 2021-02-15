package _structures

type ListNode struct {
	Val  interface{}
	Prev *ListNode
	Next *ListNode
}

type List struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func MakeList() List {
	l := List{head: &ListNode{}, tail: nil, length: 0}
	l.tail = l.head
	return l
}

func (this *List) Add(pos int, elem interface{}) bool {
	if pos > this.length {
		return false
	}
	p := this.head
	for i := 0; i < pos; i++ {
		p = p.Next
	}
	newNode := &ListNode{elem, p, p.Next}
	if p.Next == nil {
		this.tail = newNode
	} else {
		p.Next.Prev = newNode
	}
	p.Next = newNode
	this.length++
	return true
}

func (this *List) AddBack(elem interface{}) {
	this.tail.Next = &ListNode{elem, this.tail, nil}
	this.tail = this.tail.Next
	this.length++
}

func (this *List) Remove(elem interface{}) {
	prev, p := this.head, this.head.Next
	for p != nil {
		if p.Val == elem {
			p.Next.Prev = prev
			prev.Next = p.Next
		} else {
			prev = p
		}
		p = p.Next
	}
}

func (this *List) RemoveBack() interface{} {
	if this.IsEmpty() {
		return nil
	}
	res := this.tail.Val
	this.tail = this.tail.Prev
	this.tail.Next = nil
	return res
}

func (this *List) GetElem(pos int) interface{} {
	if pos >= this.length || pos < 0 {
		return nil
	}
	var p *ListNode

	if pos < this.length/2 {
		p = this.head
		for i := 0; i < pos; i++ {
			p = p.Next
		}
	} else {
		p = this.tail
		pos = this.length - pos
		for i := 0; i < pos; i++ {
			p = p.Prev
		}
	}
	return p.Next.Val
}

func (this *List) GetLastElem() interface{} {
	return this.tail.Val
}

func (this *List) IsEmpty() bool {
	return this.length == 0
}

func (this *List) Traverse(handle func(interface{})) {
	p := this.head.Next
	for p != nil {
		handle(p.Val)
		p = p.Next
	}
}

func TListTest() {
	l := MakeList()
	for i := 0; i < 10; i++ {
		l.AddBack(i)
	}
	l.Traverse(func(i interface{}) {
		println(i.(int))
	})
}
