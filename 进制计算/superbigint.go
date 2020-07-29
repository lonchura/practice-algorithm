package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	a := "1123432"
	b := "56438"

	ab := mul(a, b)
	fmt.Printf("%s * %s = %s\n", a, b, ab)
}

func mul(a string, b string) string {
	sum := "0"
	bLen := len(b)
	for i:=bLen-1; i>=0; i-- {
		// 计算一部分
		part := mulNormal(a, b[i:i+1])
		// 结尾补0
		part = part + strings.Repeat("0", bLen-1 - i)
		// 每部分累加
		sum = add(sum, part)
	}

	return sum
}

func add(a string, b string) string {
	aLen, bLen := len(a), len(b)
	maxLen := int(math.Max(float64(aLen), float64(bLen)))
	addOne := false
	result := ""
	aInt,bInt := 0, 0
	for i:=0; i<maxLen; i++ {
		// 位计算
		if i>=aLen {
			aInt = 0
			bInt, _ = strconv.Atoi(b[bLen-i-1:bLen-i])
		} else if i>=bLen {
			aInt, _ = strconv.Atoi(a[aLen-i-1:aLen-i])
			bInt = 0
		} else {
			aInt, _ = strconv.Atoi(a[aLen-i-1:aLen-i])
			bInt, _ = strconv.Atoi(b[bLen-i-1:bLen-i])
		}
		rInt := aInt + bInt
		// 进位计算
		if addOne {
			rInt += 1
		}
		// 进位判断
		if rInt >= 10 {
			addOne = true
		} else {
			addOne = false
		}
		// 位追加
		if addOne {
			result = strconv.Itoa(rInt % 10) + result
		} else {
			result = strconv.Itoa(rInt) + result
		}
	}

	return result
}

func mulNormal(a string, char string) string  {
	if char == "0" {
		return "0"
	} else if char == "1" {
		return a
	} else {
		aLen := len(a)
		result := ""
		addOne := false // 是否进位
		addNum := 0 // 进位数量
		for i:=aLen-1; i>=0; i-- {
			// 位计算
			aInt, _ := strconv.Atoi(a[i:i+1])
			cInt, _ := strconv.Atoi(char)
			rInt := aInt * cInt
			// 进位计算
			if addOne {
				rInt += addNum
			}
			// 进位判断
			if rInt >= 10 {
				addOne = true
				addNum = rInt / 10
			} else {
				addOne = false
			}
			// 追加位
			if addOne {
				result = strconv.Itoa(rInt % 10) + result
			} else {
				result = strconv.Itoa(rInt) + result
			}
		}
		return result
	}
}