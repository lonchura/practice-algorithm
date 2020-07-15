package main

import "fmt"

func main()  {
	str := "abccdbc"
	patter := "bcc"

	result := match(patter, str)
	for _,part := range result {
		fmt.Printf("%s\n", part)
	}
}

func match(patter string, str string) []string {
	_len := len(str)
	_pLen := len(patter)
	letter := patter[0:1]

	_matchs := []string{}
	for i:=0; i<_len; i++ {
		if letter == str[i:i+1] {
			if _pLen == 1 {
				// 递归结束
				_matchs = append(_matchs, fmt.Sprintf("%s", str[0:i+1]))
			} else {
				_list := match(patter[1:], str[i+1:])
				for _, part := range _list {
					// 拼接
					_matchs = append(_matchs, fmt.Sprintf("%s%s", str[0:i+1], part))
				}
			}
		}
	}

	return _matchs
}

