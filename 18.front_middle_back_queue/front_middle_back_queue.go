package main

import "fmt"

type FrontMiddleBackQueue struct {
	queue []int
}

func Constructor() FrontMiddleBackQueue {
	queue := []int{}
	return FrontMiddleBackQueue{queue}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.queue = append([]int{val}, this.queue...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	i := len(this.queue) / 2
	this.queue = append(this.queue, 0)
	copy(this.queue[i+1:], this.queue[i:])
	this.queue[i] = val
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.queue = append(this.queue, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.queue) == 0 {
		return -1
	}
	result := this.queue[0]
	this.queue = this.queue[1:]
	return result
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.queue) == 0 {
		return -1
	}
	var result int
	i := len(this.queue) / 2
	if len(this.queue)%2 != 0 {
		result = this.queue[i]
		this.queue = append(this.queue[:i], this.queue[i+1:]...)
	} else {
		result = this.queue[i-1]
		this.queue = append(this.queue[:i-1], this.queue[i:]...)
	}
	return result
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.queue) == 0 {
		return -1
	}
	result := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return result
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
func main() {
	Q := Constructor()
	Q.PushMiddle(6)
	fmt.Println(Q)
	fmt.Println(Q.PopFront())
}
