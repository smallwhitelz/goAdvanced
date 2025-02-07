package main

import "fmt"

func main() {
	// channel的定义
	// var关键字 变量名 = make(chan 类型, 长度)
	var ch1 = make(chan int, 3)
	var ch2 = make(chan bool, 5)
	var ch3 = make(chan string, 3)
	fmt.Println(ch1, ch2, ch3)
	// channel操作
	// 在channel中写入元素
	ch1 <- 10
	ch1 <- 9
	ch1 <- 8
	// 当channel长度满了时，再继续写入会死锁
	//ch1 <- 7
	// 从channel中取出元素
	num1 := <-ch1
	num2 := <-ch1
	num3 := <-ch1
	// 当channel中没有元素时，从channel中取出会死锁
	//num4 := <-ch1
	fmt.Println(num1, num2, num3)

	// 优雅的从channel中取值
	ch4 := make(chan int, 5)
	ch4 <- 10
	ch4 <- 9
	ch4 <- 8
	ch4 <- 7
	ch4 <- 6
	close(ch4)
	// 通道关闭后会退出for range循环，且不会死锁
	for i := range ch4 {
		fmt.Println(i)
	}
	// 通道关闭后，从channel中取值会返回类型零值，且不会死锁
	num4 := <-ch4
	fmt.Println(num4)

}

func Run(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			{
				fmt.Println("stop")
			}
		}
	}
}
