package main

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// 声明一个channel
	//var ch chan struct{}
	// 声明并初始化一个channel
	//ch1 := make(chan int)
	// 这种是带 buffer 的
	ch2 := make(chan int, 2)
	// 把 123 发送到ch2 里面
	ch2 <- 123
	data := <-ch2
	t.Log(data)
	// 关闭channel
	close(ch2)
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	// ok 代表有没有读到
	val, ok := <-ch
	t.Log("读到数据了吗?", ok, val)
	close(ch)
	// 这个操作会引起panic
	//ch <-123
	val, ok = <-ch
	t.Log("读到数据了吗?", ok, val)
	// 也会panic
	//close(ch)
}

func TestChannelLoop(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for val := range ch {
		t.Log(val)
	}
}

func TestChannelBlocking(t *testing.T) {
	ch := make(chan int)
	b1 := BigStruct{}
	go func() {
		var b BigStruct
		// 这个就是 goroutine 泄漏
		ch <- 123
		t.Log(b, b1)
	}()
}

type BigStruct struct {
}

func TestChannelSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	go func() {
		time.Sleep(time.Second)
		ch1 <- 123
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- 123
	}()
	select {
	case val := <-ch1:
		t.Log("进来了 ch1 这里", val)
	case val := <-ch2:
		t.Log("进来了 ch2 这里", val)
	}
}
