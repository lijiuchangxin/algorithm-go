package main

import (
	"fmt"
)

// 定义猫的结构体
type CatNode struct {
	No int			// 猫的编号
	Name string
	Next *CatNode
}

func InsertCatNode(head *CatNode, NewCatNode *CatNode)  {

	//  判断是不是添加第一只猫
	if head.Next == nil {
		head.No = NewCatNode.No
		head.Name = NewCatNode.Name
		head.Next = head // 形成一个环状
		fmt.Println(NewCatNode, "入链")
		return
	}

	// 定一个临时变量，帮忙寻找环形的最后节点
	temp := head
	for {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}

	// 加入链表中
	temp.Next = NewCatNode
	NewCatNode.Next = head
}


// 删除一只猫
// 删除一个环形但链表的思路如下：
// 1. 先让temp指向head
// 2. 让helper指向链表最后
// 3. 让temp和要删除的id比较， 如果相同，则同helper完成删除，【必须考虑删除头节点的情况】
func DeleteCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	// 空链表情况
	if temp.Next == nil {
		fmt.Println("empty circle link, not delete")
		return head
	}

	// 如果只有一个节点
	if temp.Next == head { // 说明只有一个节点
		if temp.No == id {
			temp.Next = nil
		}
		return head
	}

	// 将helper定位到链表最后
	for {
		if helper.Next == head {
			break
		}
		helper = helper.Next
	}


	// 如果有两以上包含两个节点
	flag := true
	for {
		if temp.Next == head { // 说明比较至最后一个，还没比较
			break
		}
		if temp.No == id {
			if temp == head { // 说明删除的是头节点
				head = head.Next
			}
			// 找到，我们也可以直接删除
			helper.Next = temp.Next
			//temp.Next = helper.Next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.Next    // 移动 用于比较
		helper = temp.Next  // 移动 用于删除
	}

	if flag { // 如果flag 为真，则我们在上面没有删除
		if temp.No == id {
			helper.Next = temp.Next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Println("sorry, no cat", id)
		}
	}
	fmt.Println(head)
	return head
}

// 输出环形链表
func ListCircleLink(head *CatNode)  {
	temp := head
	if temp.Next == nil {
		fmt.Println("empty circle link")
		return
	}
	for {
		fmt.Printf("cat info =[id=%d name=%s]->\n", temp.No, temp.Name)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
}

func main() {
	// 初始化一个环形链表的头节点
	head := &CatNode{}

	// 创建一只猫
	cat1 := &CatNode{
		No:   1,
		Name: "pk",
	}

	cat2 := &CatNode{
		No:   2,
		Name: "lzs",
	}

	cat3 := &CatNode{
		No:   3,
		Name: "lsy",
	}

	cat4 := &CatNode{
		No:   4,
		Name: "mjq",
	}
	//
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	InsertCatNode(head, cat4)
	ListCircleLink(head)

	head = DeleteCatNode(head, 2)
	fmt.Println(head)
	fmt.Println(head.Next)
	fmt.Println(head.Next.Next)
	fmt.Println(head.Next.Next.Next)

	//ListCircleLink(head)

}