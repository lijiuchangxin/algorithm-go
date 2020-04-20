package main

import (
	"fmt"
)

type MyCircularQueue struct {
	MaxSize int
	Head int
	Tail int
	ValueSlice []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	queue := MyCircularQueue{
		MaxSize:    k+1,
		Head:       0,
		Tail:       0,
		ValueSlice: make([]int, k+1, k+1),
	}
	return queue
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() { return false }
	this.ValueSlice[this.Tail] = value
	this.Tail =  (this.Tail + 1) % this.MaxSize
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() { return false }
	this.Head = (this.Head + 1) % this.MaxSize
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() { return -1 }
	return this.ValueSlice[this.Head]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() { return -1 }
	if this.Tail == 0 { return this.ValueSlice[this.MaxSize-1] }
	return this.ValueSlice[this.Tail-1]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return	(this.Tail + 1) % this.MaxSize == this.Head
}

func main() {
	//queue := Constructor(3)
	//fmt.Println(queue.ValueSlice)
	//a := queue.EnQueue(2)
	//fmt.Println(queue.ValueSlice)
	//
	//b := queue.Rear()
	//c := queue.Front()
	//d := queue.DeQueue()
	//e := queue.Front()
	//f := queue.DeQueue()
	//g := queue.Front()
	//h := queue.EnQueue(4)
	//i := queue.EnQueue(2)
	//j := queue.EnQueue(2)
	//k := queue.EnQueue(3)
	//fmt.Println(a, b, c, d, e, f, g, h, i, j, k)



	queue := Constructor(3)
	a := queue.EnQueue(1)
	b := queue.EnQueue(2)
	c := queue.EnQueue(3)
	d := queue.EnQueue(4)
	fmt.Println(queue.ValueSlice)
	e := queue.Rear()
	fmt.Println(e)
	f := queue.IsFull()
	g := queue.DeQueue()
	fmt.Println(queue.ValueSlice)
	h := queue.EnQueue(4)
	fmt.Println(queue.ValueSlice)
	fmt.Println("------")
	i := queue.Rear()

	fmt.Println(a, b, c, d, e, f, g, h, i)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();+
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
