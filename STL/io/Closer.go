package main

import "os"

func main()  {
	file, err := os.Open("studygolang.txt")
	defer file.Close()
	if err != nil {
		println(err)
	}
}