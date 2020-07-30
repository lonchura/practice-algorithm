package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main()  {
	// 生成100个元素数组（范围30-60）
	arr := generateRandomArr(30, 60)

	// 从数组中找到两个元素之和为100的元素
	result, err := search2EleWithSum(arr, 100)
	if err != nil {
		log.Fatal(err)
	}

	// 打印结果
	for _, find := range result {
		fmt.Printf("%v\n", find)
	}
}

// 在整型数组中搜索两个元素之和为指定结果的元素信息
// @param arr 输入的整型数组
// @param sum 两个元素之和
// @return [][]int 查找结果（结果为二维数组，每行为一个查找结果，按顺序0为第一个元素的下标值，1为第二个元素的下标值，2为第一个元素的值，3为第二个元素的值）
// @return error 错误信息
func search2EleWithSum(arr []int, sum int) ([][]int, error) {
	if len(arr) <= 1 {
		return nil, errors.New("array len need >= 2")
	}

	var result [][]int
	for i:=0; i<len(arr); i++ {
		// 查找另一个元素的值
		need := sum - arr[i]
		// 从数组中查找
		fIndex, isFound := arrIntSearch(arr, need)
		// 打印结果
		if isFound {
			// 跳过已经查找过的记录
			// 跳过指向同一个元素的记录
			if i >= fIndex {
				continue
			}
			// 追加结果
			result = append(result, []int{i, fIndex, arr[i], need})
		} else {
			continue
		}
	}

	return result, nil
}

// 范围整型数组生成
// @param from 起点值
// @param to 终点值
// @return 生成的数组（数组生成失败返回nil）
func generateRandomArr(from int, to int) []int {
	if from > to {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 100)
	ele := 0
	for i:=0; i<100; i++ {
		ele = rand.Intn(to-from) + from
		arr[i] = ele
	}

	return arr
}

// 整型数组查找
// @param arr 查找的数组
// @param need 待查找的元素值
// @return index 查找的结果索引（-1为未查找到）
// @return isFound 是否查找到
func arrIntSearch(arr []int, need int) (index int, isFound bool) {
	for i,v := range arr {
		if v == need {
			index = i
			isFound = true
			return
		}
	}

	index = -1
	isFound = false
	return
}