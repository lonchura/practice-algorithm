package main

import "fmt"

func main() {
	list := []string{"A0","A1","B0","B1","C0","C1","D0","D1"}

	result := xSort(list)
	for _,arr := range result {
		for _,val := range arr {
			fmt.Printf("%s\t", val)
		}
		fmt.Println()
	}
}

func xSort(list []string) [][]string {
	// 全排序
	_list := permutationSort(list)

	// 过滤
	for _,arr := range _list {

	}
	return _list
}

func permutationSort(list []string) [][]string {
	var result [][]string

	// list只有一个元素时，全排列为自己
	if len(list) == 1 {
		result = [][]string{list}
	} else {
		// 递归全排列
		cpList := make([]string, len(list))

		for i,first := range list {
			// new list without str
			copy(cpList, list)
			newList := append(cpList[:i], cpList[i+1:]...)

			// new list permutationSort
			partResult := permutationSort(newList)

			// result
			for _,_list := range partResult {
				result = append(result, append([]string{first}, _list...))
			}
		}
	}

	return result
}

func searchIndex(list []string, val) int {

}