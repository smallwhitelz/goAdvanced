package main

import "fmt"

// writeData
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		// 写入数据
		intChan <- i
		fmt.Println("write data:", i)
	}
	close(intChan)
}

// readData
func readData(intChan chan int, exitChan chan bool) {
	for {
		// 优雅的从intChan中取值
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read data:", v)
	}
	// 读完后，发送退出信号
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	// 异步写
	go writeData(intChan)
	// 异步读
	go readData(intChan, exitChan)
	// 监听退出信号，优雅退出
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
