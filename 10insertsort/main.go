package main

import "fmt"

func InsertSort(arr []int)  {
	// 完成第一次，给第二个元素找到合适位置并插入
	for i := 1; i < len(arr); i++ {
		insertIndex := i - 1
		insertValue := arr[i]

		// 后移
		for insertIndex >= 0 && arr[insertIndex] < insertValue {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}

		// 插入
		if insertIndex + 1 != i {
			arr[insertIndex+1] = insertValue
		}
		//arr[insertIndex+1] = insertValue
	}
}


func main() {
	//arr := []int{23, 0, 12, 56, 34}
	arr := []int{23, 0, 12, 56, 34, 66, 55 ,7, 23, 34}
	fmt.Println(arr)
	InsertSort(arr)
	fmt.Println(arr)
}
