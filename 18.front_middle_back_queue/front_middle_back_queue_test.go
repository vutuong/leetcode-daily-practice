package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	q := Constructor()
	q.PushFront(1) // [1]
	if !reflect.DeepEqual(q.queue, []int{1}) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", q.queue, []int{1})
	}
	q.PushBack(2) // [1, 2]
	if !reflect.DeepEqual(q.queue, []int{1, 2}) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", q.queue, []int{1, 2})
	}
	q.PushMiddle(3) // [1, 3, 2]
	if !reflect.DeepEqual(q.queue, []int{1, 3, 2}) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", q.queue, []int{1, 3, 2})
	}
	q.PushMiddle(4) // [1, 4, 3, 2]
	if !reflect.DeepEqual(q.queue, []int{1, 4, 3, 2}) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", q.queue, []int{1, 4, 3, 2})
	}
	test1 := q.PopFront() // return 1 -> [4, 3, 2]
	if !reflect.DeepEqual(q.queue, []int{4, 3, 2}) && test1 != 1 {
		t.Errorf("Result was incorrect, got: %v, %v, want: %v, %v.", q.queue, test1, []int{4, 3, 2}, 1)
	}
	test2 := q.PopMiddle() // return 3 -> [4, 2]
	if !reflect.DeepEqual(q.queue, []int{4, 2}) && test2 != 3 {
		t.Errorf("Result was incorrect, got: %v, %v, want: %v, %v.", q.queue, test2, []int{4, 2}, 3)
	}
	test3 := q.PopMiddle() // return 4 -> [2]
	if !reflect.DeepEqual(q.queue, []int{2}) && test3 != 4 {
		t.Errorf("Result was incorrect, got: %v, %v, want: %v, %v.", q.queue, test3, []int{2}, 4)
	}
	test4 := q.PopBack() // return 2 -> []
	if !reflect.DeepEqual(q.queue, []int{}) && test4 != 2 {
		t.Errorf("Result was incorrect, got: %v, %v, want: %v, %v.", q.queue, test4, []int{}, 2)
	}
	test5 := q.PopFront() // return -1 -> [] (The queue is empty)
	if !reflect.DeepEqual(q.queue, []int{}) && test5 != -1 {
		t.Errorf("Result was incorrect, got: %v, %v, want: %v, %v.", q.queue, test5, []int{}, -1)
	}
}
