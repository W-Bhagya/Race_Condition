package main

import (
	"fmt"
	"sync/atomic"
	"time"

)

var num int64 = 0

func main() {
	

	for i := 0; i < 100; i++ {
		go increment()
	}

     time.Sleep(time.Second * 2)

	fmt.Println(num)
}
func increment() {
	num++
	atomic.AddInt64(&num, 1)
}