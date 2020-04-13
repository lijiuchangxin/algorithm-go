package main

import (
	"fmt"
)


// 定义一个heronode
type HeroNode struct {
	No 			int
	Name 		string
	NickName	string
	Next		*HeroNode   // 表示指向下一个节点
}

// 给链表插入节点
// 第一种插入方式，在单链表最后加入，
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode)  {
	// 思路
	// 1. 先找到该链表的最后节点
	// 2. 创建一个辅助节点
	temp := head
	for {
		if temp.Next == nil { // 表示找到最后
			 break
		}
		temp = temp.Next // 让temp不断的指向下一个节点
	}
	// 3. 将newHeroNode加入到链表的最后
	temp.Next = newHeroNode
}

// 显示链表的所有的节点信息
func ListHeroNode(head *HeroNode) {
	// 1. 创建一个辅助节点
	temp := head

	// 先判断该链表是不是一个空链表
	if temp.Next == nil {
		fmt.Println("空空如也")
		return
	}
	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s, %s]==>", temp.Next.No, temp.Next.Name, temp.Next.NickName )
		// 判断是否在链表最后
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
}

func main() {

	// 1. 创建一个头节点
	head := &HeroNode{}

	// 2. 创建一个新的heronode
	hero1 := &HeroNode{1, "宋江", "及时雨", nil}
	hero2 := &HeroNode{2, "李逵", "黑旋风", nil}
	hero3 := &HeroNode{3, "卢俊义", "玉麒麟", nil}

	// 3. 加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)

	// 4. 显示
	ListHeroNode(head)
}