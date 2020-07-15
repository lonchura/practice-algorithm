package main

import (
	"fmt"
	"strings"
)

func main() {
	//str1 := "123sdf"
	str1 := "12"
	//str2 := "45dr"
	str2 := "12"

	str3 := addition36(str1, str2)

	fmt.Print(str3)
}

func addition36(str1 string, str2 string) (string3 string) {
	int1 := convert32to10(str1)
	int2 := convert32to10(str2)
	fmt.Println(int1)
	fmt.Println(int2)

	int3 := int1 + int2
	string3 = convert10to32(int3)

	return
}

func convert32to10(str string) int {
	map32 := "0123456789abcdefghijklmnopqrstuvwxyz"

	sLen := len(str)
	sum := 0
	times := 1
	for i:=sLen-1; i>=0; i-- {
		letter := str[i:i+1]
		val := strings.Index(map32, letter)

		sum += val * times
		times *= 36
	}

	return sum
}

func convert10to32(num int) string {
	map32 := "0123456789abcdefghijklmnopqrstuvwxyz"

	str := ""
	for {
		i := num % 36
		letter := map32[i:i+1]
		str = fmt.Sprintf("%s%s", letter, str)

		if num < 36 {
			break
		}
		num = num / 36
	}


	return str
}