package main

import (
	"bytes"
	"os"
)

func main()  {
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}