package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup

// 互斥锁
var mtx sync.Mutex

// 累加
func add() {
	defer wg.Done()
	mtx.Lock()
	num += 1
	mtx.Unlock()
}

func main() {
	// 开启1000个协程，执行add函数
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(num)
}
