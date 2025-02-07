package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test方法，你好golang")
		time.Sleep(time.Millisecond * 100)
	}
}

// main就是一个主线程，在主线程中使用go关键字开启一个协程运行test方法
// 这种情况下，当main执行完毕退出后，协程也会退出
func main() {
	// 启动goroutine
	go test()
	for i := 0; i < 2; i++ {
		fmt.Println("main方法，你好golang")
		time.Sleep(time.Millisecond * 100)
	}

}
