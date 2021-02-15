/*
 * https://leetcode-cn.com/problems/exam-room/
 * 2020.11.20
 * Golang 16ms(72.73%), 6.6MB(72.73%)
 */
package design

import "container/list"

type ExamRoom struct {
	students *list.List
	size     int
}

func ExamRoomConstructor(N int) ExamRoom {
	return ExamRoom{list.New(), N - 1}
}

func (this *ExamRoom) Seat() int {
	if this.students.Len() == 0 {
		this.students.PushFront(0)
		return 0
	}

	p := this.students.Front()
	pre := p.Value.(int)
	max, addVal, addFront := pre, 0, p

	for p = p.Next(); p != nil; p = p.Next() {
		val := p.Value.(int)
		distance := (p.Value.(int) - pre) / 2
		if distance > max {
			max = distance
			addFront = p
			addVal = pre + distance
		}
		pre = val
	}

	distance := this.size - pre
	if distance > max {
		this.students.PushBack(this.size)
		return this.size
	}
	this.students.InsertBefore(addVal, addFront)
	return addVal
}

func (this *ExamRoom) Leave(p int) {
	for e := this.students.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			this.students.Remove(e)
			return
		}
	}
}
