package main

type Node struct {
	Value int
	Next *Node
}

type MyStack struct {
	Front *Node
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	mystack := MyStack{}
	return mystack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	node := &Node{
		Value: x,
		Next:  this.Front,
	}

	this.Front = node
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Front != nil {
		res := this.Front.Value
		this.Front.Next, this.Front = nil, this.Front.Next
		return res
	}
	return -1
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Front == nil { return -1 }
	return this.Front.Value
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Front == nil
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */