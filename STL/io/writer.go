package main

import "fmt"

func main()  {
	ret, err := fmt.Printf("Hello %s!\n", "こんにちは")
	fmt.Println(ret, err)
}