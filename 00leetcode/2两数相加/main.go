package main

import (
	"fmt"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//Line 115: Char 37: cannot use param_1 (type *leetcode.ListNode) as type *ListNode in argument to __helper__ (solution.go)

type ListNode struct {
	Val int
	Next *ListNode
}

func RetNumCount(this *ListNode) (res float64) {
	LinkLen  := CountLinkLen(this)
	var val float64
	temp := this
	for i := LinkLen; i > 0; i-- {
		val = math.Pow(10, LinkLen-1) * float64(temp.Val)
		res += val
		LinkLen -= 1
		temp = temp.Next
	}
	return res
}

func CountLinkLen(this *ListNode) (res float64) {
	temp := this
	for {
		if temp.Next == nil { break }
		temp = temp.Next
		res += 1
	}
	return res+1
}

func RetTwoLinkLen(l1 *ListNode, l2 *ListNode) (len1, len2 float64) {
	len1, len2 = CountLinkLen(l1), CountLinkLen(l2)
	return len1, len2
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := RetNumCount(l1) + RetNumCount(l2)
	len1, len2 := RetTwoLinkLen(l1, l2)

	var len0 float64
	if len1 > len2 {
		len0 = len1+1
	} else {
		len0 = len2+1
	}

	head := &ListNode{}
	temp := head

	for i := 0.0; i < len0; i++ {
		newNode := &ListNode{}
		val := float64(int(res/math.Pow(10, i))%10)
		if val == 0 && i == len0 -1 { break }
		temp.Next = newNode
		//if i <= len0-1 { temp.Next = newNode }
		temp.Val = int(val)
		fmt.Println(temp.Val)
		temp = newNode
	}
	return head
}

func ShowLink(li *ListNode)  {
	temp := li
	for {
		if temp.Next == nil {
			break
		}
		fmt.Printf("%d =>", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func main() {
	//243
	//564
	n10 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n11 := &ListNode{
		Val:  4,
		Next: nil,
	}
	n12 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n10.Next = n11
	n11.Next = n12
	//fmt.Println(	n10.RetNumCount())
	//fmt.Println(	n10.CountLinkLen())
	//
	n20 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n21 := &ListNode{
		Val:  6,
		Next: nil,
	}
	n22 := &ListNode{
		Val:  4,
		Next: nil,
	}
	n20.Next = n21
	n21.Next = n22
	n := addTwoNumbers(n10, n20)
	ShowLink(n)



	//5
	//5
	a := &ListNode{5, nil}
	b := &ListNode{5, nil}
	ShowLink(addTwoNumbers(a, b))

}