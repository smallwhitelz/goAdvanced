package main

import "fmt"

func main() {
	// intChan
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// stringChan
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("%d", i)
	}
	// select监听两个管道
	for {
		select {
		case v := <-intChan:
			fmt.Println("intChan:", v)
		case v := <-stringChan:
			fmt.Println("stringChan:", v)
		default:
			fmt.Println("default")
			return
		}

	}
}
