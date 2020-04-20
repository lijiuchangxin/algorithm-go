package main

import "fmt"

// 快速排序
// 说明
// 1. left  表示数组左边的下标
// 2. right 表示数组右边的下标
// 3. array 表示要排序的数组
func QuickSort(left int, right int, array []int)  {
	l := left
	r := right

	// pivot 中轴
	pivot := array[(left + right) / 2]

	// for循环的目标是将比pivot小的数放到左边
	// 比pivot大的数放到右边
	for ; l < r; {
		// 从pivot的左边找到大于等于pivot的值
		for ; array[l] < pivot; {
			l++
		}
		// 从pivot的右边找到小于等于pivot的值
		for ; array[r] > pivot; {
			r--
		}
		// 如果 l>=r 表明本次分解任务完成
		if l >= r {
			break
		}
		// 交换
		array[l], array[r] = array[r], array[l]

		// 优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}

		// 向左递归
		if left < r {
			QuickSort(left, r, array)
		}
		// 向右递归
		if right > l {
			QuickSort(l, right, array)
		}
	}

}


func main() {

	//arr := []int{-9, 78, 0, 23, -567, 70, -12, 45, 99, 91, 832}
	arr := []int{-9, 78, 0, 23, -567, 70, -7, 0, 99, 1203, -78, 884, -455}
	fmt.Println(arr)
	QuickSort(0, len(arr)-1, arr)
	fmt.Println(arr)
}
