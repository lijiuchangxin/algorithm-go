package main

import "fmt"

// 编写函数selectShort 完成排序

func SelectSort(arr []int)  {

	for j := 0; j < len(arr) -1; j++ {
		// 1. 先完成将第一个大值和 arr[0] => 先易后难
		max := arr[j]
		maxIndex := j

		// 2.遍历后面1-[len(arr)-1] 比较
		for i := j+1; i < len(arr); i++ {
			if max < arr[i] { //  找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}

		// 交换
		if maxIndex != 0 {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}



}


func main() {
	// 定一个数组
	arr := []int{20, 34, 19, 100, 80, 789, 60, 60, 62}
	SelectSort(arr)
	fmt.Println(arr)
}
