package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}


func main() {
	// 1. 创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1  // 黑子
	chessMap[2][3] = 2  // 蓝子

	// 2. 输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组
	// 思路
	// (1). 遍历chessMap，如果我们发现有一个元素的值不为0，创建一个node结构体
	// (2). 将其放入对应的切片即可

	var sparseArr []ValNode

	// 标准的稀疏数组应该还有一个 记录元素的二维数组的规模（行列默认值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个valnode 值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	// 将稀疏数组，存盘

	// 如何恢复原始数组
	// 1. 创建一个原始数组
	var chessMap2 [11][11]int

	// 2. 遍历 sparseArr
	for i, valueNode := range sparseArr {
		if i != 0 {
			chessMap2[valueNode.row][valueNode.col] = valueNode.val
		}
	}

	fmt.Println("recover.....")
	// 3. 看看chessMap2是不是回复
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}}
