package main

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

var wg sync.WaitGroup
var finished = make(chan int)

func main() {
	pipeReader, pipeWriter := io.Pipe()
	wg.Add(1)
	go PipeWrite(pipeWriter)
	go PipeRead(pipeReader)
	wg.Wait()
}

func PipeWrite(writer *io.PipeWriter){
	data := []byte("Go语言中文网")
	for i := 0; i < 3; i++{
		n, err := writer.Write(data)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("写入字节 %d\n",n)
		finished <- 0
	}
	writer.CloseWithError(errors.New("写入段已关闭"))
	finished <- 1
}

func PipeRead(reader *io.PipeReader){
	var bool int
	buf := make([]byte, 128)
	for{
		fmt.Println("接口端开始阻塞5秒钟...")
		time.Sleep(1 * time.Second)
		fmt.Println("接收端开始接受")
		n, err := reader.Read(buf)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("收到字节: %d\n buf内容: %s\n",n,buf)

		bool = <- finished
		fmt.Println(bool)
		if bool==1 {
			fmt.Println("Done!!!")
			wg.Done()
		} else {
			fmt.Println("Doing...")
		}
	}
}