package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main()  {
	// share memory int
	var i int32 = 0

	wg := sync.WaitGroup{}
	//mutex := sync.Mutex{}


	wg.Add(1)
	go func() {
		for m:=0; m<1000000; m++ {
			//mutex.Lock()
			//i++
			//mutex.Unlock()

			atomic.AddInt32(&i, 1)
		}
		wg.Done()
		fmt.Printf("goroutine 1 done!\n")
	}()
	wg.Add(1)
	go func() {
		for n:=0; n<1000000; n++ {
			//mutex.Lock()
			//i++
			//mutex.Unlock()

			atomic.AddInt32(&i, 1)
		}
		wg.Done()
		fmt.Printf("goroutine 2 done!\n")
	}()
	wg.Add(1)
	go func() {
		for n:=0; n<1000000; n++ {
			//mutex.Lock()
			//i++
			//mutex.Unlock()

			atomic.AddInt32(&i, 1)
		}
		wg.Done()
		fmt.Printf("goroutine 3 done!\n")
	}()
	wg.Add(1)
	go func() {
		for n:=0; n<1000000; n++ {
			//mutex.Lock()
			//i++
			//mutex.Unlock()

			atomic.AddInt32(&i, 1)
		}
		wg.Done()
		fmt.Printf("goroutine 4 done!\n")
	}()
	wg.Add(1)
	go func() {
		for n:=0; n<1000000; n++ {
			//mutex.Lock()
			//i++
			//mutex.Unlock()

			atomic.AddInt32(&i, 1)
		}
		wg.Done()
		fmt.Printf("goroutine 5 done!\n")
	}()
	wg.Wait()

	fmt.Printf("goroutine i++ done, i:%d\n", i)
}