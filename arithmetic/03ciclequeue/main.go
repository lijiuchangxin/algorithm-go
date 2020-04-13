package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列
type circleQueue struct {
	maxSize int //5
	array [5]int
	head int  // 指向队列对首
	tail int  // 指向队尾
}

// 入队列
func (this *circleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}

	// 分析出this.tail 在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val  // 把值给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 出队列
func (this *circleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}

	// 取出, head指向对首，并且包含对首的元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, nil
}

// 判断环形队列满了
func (this *circleQueue) IsFull() bool {
	return (this.tail + 1) % this.maxSize == this.head  // ？？？
}

// 判断环形队列是否为空
func (this *circleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 判断环形队列有多少个元素
func (this *circleQueue) Size() int {
	// 这是一个关键算法
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func (this *circleQueue) ListQueue() {

	fmt.Println("环形队列情况如下：")
	// 取出当前队列有多少元素
	size := this.Size()
	if size == 0 {
		fmt.Println("queue empty")
	}

	// 设计一个辅助变量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead+1)%this.maxSize
	}
	fmt.Println()

}

func main()  {

	queue := &circleQueue{
		maxSize: 5,
		array:   [5]int{},
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列中获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队到列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数：", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
			}
	}
}

