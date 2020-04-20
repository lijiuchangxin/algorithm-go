package main

import (
	"fmt"
)

type Node struct {
	Val 	int
	Next 	*Node
	Pre 	*Node
}


type MyLinkedList struct {
	Head *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := &Node{}
	return MyLinkedList{head}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	linkLen := this.CountLinkLen()
	nowIndex := 0
	temp := this.Head
	if temp.Next == nil || linkLen-1 < index {  // 空链表或者链表长度小于索引
		return -1
	}

	for {
		if nowIndex == index {
			return temp.Next.Val
		}
		temp = temp.Next
		nowIndex += 1
	}
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	newNode := &Node{Val:val}
	if this.Head.Next == nil { // 如果只是空链表，头节点后无节点
		this.Head.Next = newNode
		newNode.Pre = this.Head
	} else {
		newNode.Next = this.Head.Next
		this.Head.Next.Pre = newNode
		this.Head.Next = newNode
		newNode.Pre = this.Head
	}
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	temp := this.Head
	newNode := &Node{Val:val}
	for {
		if temp.Next == nil {
			temp.Next = newNode
			newNode.Pre = temp
			break
		}
		temp = temp.Next
	}
}

func (this *MyLinkedList)CountLinkLen() int {
	nowLen := 0
	temp := this.Head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		nowLen += 1
	}
	return nowLen
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	linkLen := this.CountLinkLen()
	if linkLen < index {  // 如果索引超过link长度不添加
		return
	} else if linkLen == index {  // 如果索引等于link长度追加到末尾
		this.AddAtTail(val)
	} else {
		newNode := &Node{Val:val}
		temp := this.Head
		nowIndex := 0
		for {
			if nowIndex == index {
				newNode.Next = temp.Next
				temp.Next.Pre = newNode
				temp.Next = newNode
				newNode.Pre = temp
				break
			}
			nowIndex += 1
			temp = temp.Next
		}
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	linkLen := this.CountLinkLen()
	if linkLen -1 <= index {
		return
	}

	nowIndex := 0
	temp := this.Head
	for {
		if nowIndex == index {
			if temp.Next.Next == nil {
				temp.Next = nil
				break
			}
			temp.Next.Next.Pre = temp
			temp.Next = temp.Next.Next
			break
		}
		temp = temp.Next
		nowIndex += 1
	}
}

func (this *MyLinkedList) ShowLink() {
	temp := this.Head
	index := 0
	for {
		if temp.Next == nil {
			break
		}

		fmt.Printf("%d, %d=>", index, temp.Next.Val)
		temp = temp.Next
		index += 1
	}
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	//obj := MyLinkedList{Head:&Node{}}
	//obj.AddAtHead(1)
	//obj.DeleteAtIndex(0)

	//obj := MyLinkedList{Head:&Node{}}
	//obj.AddAtHead(2)
	//obj.DeleteAtIndex(1)
	//obj.AddAtHead(2)
	//obj.AddAtHead(7)
	//obj.AddAtHead(3)
	//obj.AddAtHead(2)
	//obj.AddAtHead(5)
	//obj.AddAtTail(5)
	//obj.Get(5)
	//obj.DeleteAtIndex(6)
	//obj.DeleteAtIndex(4)

	//["MyLinkedList","AddAtHead","Get","AddAtHead","AddAtHead","DeleteAtIndex","AddAtHead","Get","Get","Get","AddAtHead","DeleteAtIndex"]
	//[[],[4],[1],[1],[5],[3],[7],[3],[3],[3],[1],[4]]
	//obj := MyLinkedList{Head:&Node{}}
	//obj.AddAtHead(4)
	//obj.Get(1)
	//obj.AddAtHead(1)
	//obj.AddAtHead(5)
	//obj.DeleteAtIndex(3)
	//obj.AddAtHead(7)
	//obj.Get(3)
	//obj.Get(3)
	//obj.Get(3)
	//obj.AddAtHead(1)
	//obj.DeleteAtIndex(4)

	//"AddAtHead","AddAtTail","AddAtTail","AddAtHead","Get","AddAtTail","AddAtHead","AddAtTail","Get","AddAtIndex"]
	// [60],[26],[4],[9],[45],[38],[95],[78],[54],[42,86]
	//obj := MyLinkedList{Head:&Node{}}
	//obj.AddAtHead(60)
	//obj.AddAtTail(26)
	//obj.AddAtTail(4)
	//obj.AddAtHead(9)
	//b := obj.Get(3)
	//obj.AddAtTail(38)
	//obj.AddAtHead(95)
	//obj.AddAtTail(78)
	//a := obj.Get(5)
	//obj.AddAtIndex(4, 86)
	//fmt.Println(b, a)

	obj := MyLinkedList{Head:&Node{}}
	//test:=[]string{"AddAtHead","AddAtTail","AddAtTail","Get","Get","AddAtTail","AddAtIndex","AddAtHead","" +
	//	"AddAtHead","AddAtTail","AddAtTail","AddAtTail","AddAtTail","Get","AddAtHead","AddAtHead","AddAtIndex","AddAtIndex",
	//	"AddAtHead","AddAtTail","DeleteAtIndex","AddAtHead","AddAtHead","AddAtIndex","AddAtTail","Get","AddAtIndex","AddAtTail","" +
	//	"AddAtHead","AddAtHead","AddAtIndex","AddAtTail","AddAtHead","AddAtHead","Get","DeleteAtIndex","AddAtTail","AddAtTail",
	//	"AddAtHead","AddAtTail","Get","DeleteAtIndex","AddAtTail","AddAtHead","AddAtTail","DeleteAtIndex","AddAtTail","DeleteAtIndex",
	//	"AddAtIndex","DeleteAtIndex","AddAtTail","AddAtHead","AddAtIndex","AddAtHead","AddAtHead","Get","AddAtHead","Get","AddAtHead","" +
	//	"DeleteAtIndex","Get","AddAtHead","AddAtTail","Get","AddAtHead","Get","AddAtTail","Get","AddAtTail","AddAtHead","AddAtIndex",
	//	"AddAtIndex","AddAtHead","AddAtHead","DeleteAtIndex","Get","AddAtHead","AddAtIndex","AddAtTail","Get","AddAtIndex","Get","AddAtIndex",
	//	"Get","AddAtIndex","AddAtIndex","AddAtHead","AddAtHead","AddAtTail","AddAtIndex","Get","AddAtHead","AddAtTail","AddAtTail","AddAtHead","" +
	//	"Get","AddAtTail","AddAtHead","AddAtTail","Get","AddAtIndex"}
	//
	//parm := [][]int{{84},{2},{39},{3},{1},{42},{1,80},{14},{1},{53},{98},{19},{12},{2},{16},{33},{4,17},{6,8},{37},{43},{11},{80},{31},{13,23},{17},{4},{10,0},{21},{73},{22},{24,37},{14},{97},{8},{6},{17},{50},{28},{76},{79},{18},{30},{5},{9},{83},{3},{40},{26},{20,90},{30},{40},{56},{15,23},{51},{21},{26},{83},{30},{12},{8},{4},{20},{45},{10},{56},{18},{33},{2},{70},{57},{31,24},{16,92},{40},{23},{26},{1},{92},{3,78},{42},{18},{39,9},{13},{33,17},{51},{18,95},{18,33},{80},{21},{7},{17,46},{33},{60},{26},{4},{9},{45},{38},{95},{78},{54},{42,86}}
	test:=[]string{"AddAtHead","AddAtTail","AddAtTail","Get","Get","AddAtTail","AddAtIndex","AddAtHead","" +
		"AddAtHead","AddAtTail","AddAtTail","AddAtTail","AddAtTail","Get","AddAtHead","AddAtHead","AddAtIndex","AddAtIndex",
		"AddAtHead","AddAtTail","DeleteAtIndex","AddAtHead","AddAtHead","AddAtIndex","AddAtTail","Get","AddAtIndex","AddAtTail","" +
			"AddAtHead","AddAtHead","AddAtIndex","AddAtTail","AddAtHead","AddAtHead","Get","DeleteAtIndex","AddAtTail","AddAtTail",
		"AddAtHead","AddAtTail","Get","DeleteAtIndex","AddAtTail","AddAtHead","AddAtTail","DeleteAtIndex","AddAtTail","DeleteAtIndex",
		"AddAtIndex","DeleteAtIndex","AddAtTail","AddAtHead","AddAtIndex","AddAtHead","AddAtHead","Get","AddAtHead","Get","AddAtHead","" +
			"DeleteAtIndex","Get","AddAtHead","AddAtTail","Get","AddAtHead","Get","AddAtTail","Get","AddAtTail","AddAtHead","AddAtIndex",
		"AddAtIndex","AddAtHead","AddAtHead","DeleteAtIndex","Get","AddAtHead","AddAtIndex","AddAtTail","Get","AddAtIndex","Get","AddAtIndex",
		"Get","AddAtIndex","AddAtIndex","AddAtHead","AddAtHead","AddAtTail","AddAtIndex","Get","AddAtHead","AddAtTail","AddAtTail","AddAtHead","" +
			"Get","AddAtTail","AddAtHead","AddAtTail","Get"}

	parm := [][]int{{84},{2},{39},{3},{1},{42},{1,80},{14},{1},{53},{98},{19},{12},{2},{16},{33},{4,17},{6,8},{37},{43},{11},{80},{31},{13,23},{17},{4},{10,0},{21},{73},{22},{24,37},{14},{97},{8},{6},{17},{50},{28},{76},{79},{18},{30},{5},{9},{83},{3},{40},{26},{20,90},{30},{40},{56},{15,23},{51},{21},{26},{83},{30},{12},{8},{4},{20},{45},{10},{56},{18},{33},{2},{70},{57},{31,24},{16,92},{40},{23},{26},{1},{92},{3,78},{42},{18},{39,9},{13},{33,17},{51},{18,95},{18,33},{80},{21},{7},{17,46},{33},{60},{26},{4},{9},{45},{38},{95},{78},{54}}


	for index, fun := range test {
		if fun == "AddAtHead" {
			obj.AddAtHead(parm[index][0])
		}

		if fun == "Get" {
			obj.Get(parm[index][0])
		}

		if fun == "AddAtTail" {
			obj.AddAtTail(parm[index][0])
		}

		if fun == "AddAtIndex" {
			obj.AddAtIndex(parm[index][0], parm[index][1])
		}

		if fun == "DeleteAtIndex" {
			obj.DeleteAtIndex(parm[index][0])
		}

		if index == 97 {
			obj.ShowLink()
			fmt.Println()
		}

		if index == 100 {
			obj.ShowLink()
			fmt.Println()
		}


	}
}








