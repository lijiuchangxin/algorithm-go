package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理我们的队列
type Queue struct {
	maxSize int
	array [5]int // 数组模拟队列
	front int	// 表示指向队列列首
	rear int // 表示指向队列的尾部
}

// 添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {

	// 先判断队列是否已满
	if this.rear == this.maxSize - 1 {  // 重要提示：rear是队列尾部（包含）
		return errors.New("queue full")
	}

	this.rear++  // rear后移
	this.array[this.rear] = val
	return
}

// 显示队列，找到对首，然后遍历到队尾
func (this Queue) ShowQueue() {
	// this.front 不包含对首元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.rear == this.front { // 队空
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}


// 测试
func main() {
	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		array:   [5]int{},
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数：", val)
			}

		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
			}
	}
}


// 上述代码实现了基本的队列，但是数组的空间利用率很低