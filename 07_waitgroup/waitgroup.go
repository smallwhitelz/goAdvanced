package main

import (
	"fmt"
	"sync"
)

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1方法，你好golang")
		//time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2方法，你好golang")
		//time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

// WaitGroup的使用
var wg sync.WaitGroup

func main() {
	// 启动一个协程，增加1
	//wg.Add(1)
	//go test1()
	//wg.Add(1)
	//go test2()
	// 阻塞一下主线程
	//wg.Wait()
	// wg的计数器归0，就不阻塞了，会继续执行
	//fmt.Println("主线程退出")

	// 开启多个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			// defer用于函数执行完，最后执行
			// 常用场景：关闭文件句柄，关闭db连接，关闭client
			defer wg.Done()
			fmt.Println("协程", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("主线程退出")
}
