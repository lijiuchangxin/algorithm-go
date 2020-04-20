package main

import "fmt"

// 小孩结构体
type Boy struct {
	No 		int		// 编号
	Next	*Boy	// 指向下一个小孩的指针
}

// 编写一个函数，构成单项环形链表
// num 表示小孩子的个数
// *boy 返回该环形链表的第一个小孩指针
func AddBoy(num int) *Boy {

	first := &Boy{} // 空节点
	curBoy := &Boy{} // 空节点

	// 判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	// 循环的构建这个环形的链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No:   i,
		}
		// 分析构成循环链表，需要一个辅助指针【帮忙的】
		// 因为第一个小孩比较特殊
		if i == 1 { // 第一个小孩
			first = boy
			curBoy = boy
			curBoy.Next = first //
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first  // 构成环形链表
		}
	}
	return first
}


// 显示单项的环形链比【遍历】
func ShowBoy(first *Boy)  {
	// 处理环形脸变为空时
	if first.Next  == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}

	// 创建一个指针帮助遍历,[至少有一个小孩]
	curBoy := first
	for {
		fmt.Printf("小孩子编号=%d->", curBoy.No)
		// 退出条件？
		if curBoy.Next == first {
			break
		}
		// curBoy 移动到下一个
		curBoy = curBoy.Next
	}
}

/*
设编号为1，2，.... n的n个人围坐一圈，约定好编号为k，（1<=k<=n）
的人从1开始宝树，数到m的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，以此类推，知道所有的人出列位置，由此产生一个出队列
 */

// 分析思路
// 编写一个函数，PlayGame(first *Boy, startNo int, CountNum int)
// 最后我们使用一个算法，按照要求，在环形列表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int)  {
	// 1. 空链表我们单独处理
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}

	// 留一个，判断start_nun <= 小孩总数
	// 2. 需要定义辅助指针，用来帮我们删除小孩
	tail := first

	// 3. 让tail指向环形链表的最后一个小孩，这个非常重要
	// 因为tail在删除小孩的时候要用到
	for {
		if tail.Next == first { // 说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}

	// 4. 让first 移动到startNo [后面我们删除小孩，就要以first为准]
	for i := 1; i <= startNo -1; i++ {
		first = first.Next
		tail = tail.Next
	}
	// 5. 开始数countNum， 然后删除first 指向小孩
	for {
		// 开始数countNum-1次
		for i := 1; i <= countNum -1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出拳 \n", first.No)
		// 删除first指向的小孩
		first = first.Next
		tail.Next = first
		// 判断如果tail == first， 圈中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出拳 ", first.No)
}



func main() {

	first := AddBoy(5)
	// 显示
	//ShowBoy(first)
	PlayGame(first, 2, 3)
}
