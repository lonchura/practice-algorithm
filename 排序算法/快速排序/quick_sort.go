package main

import "fmt"

func main() {
	list := []int{1987, 100, 9, 250, 9, 1024, 925}
	sorted := quickSort(list)

	for _,val := range sorted {
		fmt.Printf("%d\n", val)
	}
}

func quickSort(list []int) []int {
	// 数组<2个时，无需排序直接返回
	if len(list) < 2 {
		return list
	}

	// 初始化两个数组
	// 小于等于基准值数组$left
	var leftPart []int = make([]int, 0, 0)
	// 大于基准值数组$right
	var rightPart []int = make([]int, 0, 0)

	// 选出基准值，默认选择数组第1个数据，并移除数组的第一个元素
	baseVal := list[0]

	// 循环剩余数组
	for _,val := range list[1:] {
		//fmt.Println(val)
		// 将<=基准值放入数组$left
		if val <= baseVal {
			leftPart = append(leftPart, val)
		}
		// 将>基准值放入数组$right
		if val > baseVal {
			rightPart = append(rightPart, val)
		}
	}

	// 返回 数组拼接(quickSort($left),[基准值],quickSort($right))
	return append(append(quickSort(leftPart), baseVal), quickSort(rightPart)...)
}
