package main

import (
	"fmt"
)


// 定义一个heronode
type HeroNode struct {
	No 			int
	Name 		string
	NickName	string
	Pre 		*HeroNode	// 指向前一个节点
	Next		*HeroNode   // 表示指向下一个节点
}

// 给双向链表插入节点
// 第一种插入方式，在双向链表最后加入，
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
	newHeroNode.Pre = temp
}

// 给链表插入节点
// 第二种插入方式，根据no的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode)  {
	// 思路
	// 1. 找到适当的节点位置
	// 2. 创建一个辅助节点
	temp := head
	flag := true
	// 让插入的节点no和temp的下一个节点的no进行比较
	for {
		if temp.Next == nil { // 说明到了链表的最后
			break
		} else if temp.Next.No > newHeroNode.No {
			//  说明newheronode就应该插入到temp的后面
			break
		} else if temp.Next.No == newHeroNode.No {
			// 说明链表中存在相同id的node，不让插入
			flag = false
			break
		}
		temp = temp.Next
	}
	if flag == false {
		fmt.Println("sorry, 已经存在no=", newHeroNode.No)
		return
	} else {
		newHeroNode.Next = temp.Next
		newHeroNode.Pre = temp
		if temp.Next != nil {
			temp.Next.Pre = newHeroNode
		}
		temp.Next = newHeroNode
	}
}

//	删除一个节点
func DelHeroNode(head *HeroNode, id int)  {
	temp := head
	flag := false
	// 找到要删除的节点no，和temp的下一个节点no比较
	for {
		if temp.Next == nil { // 说明到了链表的最后
			break
		} else if temp.Next.No == id {
			// 说明找到了
			flag = true
			break
		}
		temp = temp.Next
	}

	if flag { // 找到则删除
		temp.Next = temp.Next.Next
		if temp.Next != nil { temp.Next.Pre = temp }
	} else {
		fmt.Println("sorry, id 不存在")
	}
}

// 显示链表的所有的节点信息
// 仍然使用单项链表的显示方式
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

// 逆序
func ListHeroNode2(head *HeroNode) {
	// 1. 创建一个辅助节点
	temp := head

	// 先判断该链表是不是一个空链表
	if temp.Next == nil {
		fmt.Println("空空如也")
		return
	}

	// 2. 让temp定位到双向链表的最后
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s, %s]==>", temp.No, temp.Name, temp.NickName )
		// 判断是否为双链表头
		temp = temp.Pre
		if temp.Pre == nil {
			break
		}
	}
}

func main() {

	// 1. 创建一个头节点
	head := &HeroNode{}

	// 2. 创建一个新的heronode
	hero1 := &HeroNode{1, "宋江", "及时雨", nil, nil}
	hero2 := &HeroNode{2, "李逵", "黑旋风", nil, nil}
	hero3 := &HeroNode{3, "卢俊义", "玉麒麟", nil, nil}
	hero4 := &HeroNode{4, "林冲", "豹子头", nil, nil}
	hero5 := &HeroNode{5, "鲁智深", "花和尚", nil, nil}

	// 3. 加入
	InsertHeroNode2(head, hero4)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero5)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero3)
	//InsertHeroNode(head, hero1)
	//InsertHeroNode(head, hero2)
	//InsertHeroNode(head, hero3)
	//InsertHeroNode(head, hero4)
	//InsertHeroNode(head, hero5)

	// 4. 显示
	ListHeroNode(head)
	fmt.Println()
	ListHeroNode2(head)

	// 5. 删除
	DelHeroNode(head, 3)

	fmt.Println()
	DelHeroNode(head, 9)
	DelHeroNode(head, 5)
	ListHeroNode(head)

}