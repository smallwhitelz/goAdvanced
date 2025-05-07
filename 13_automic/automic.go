package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var val int32 = 12
	// 原子读，你不会读到修改一半的数据
	loadInt32 := atomic.LoadInt32(&val)
	fmt.Println(loadInt32)
	// 原子写，即便别的goroutine在别的CPU核上，也能立刻看到
	atomic.StoreInt32(&val, 13)
	// 原子自增，返回的是自增后的结果
	newVal := atomic.AddInt32(&val, 1)
	fmt.Println(newVal)
	// CAS 操作
	// 如果val的值是13，就修改为15
	swapped := atomic.CompareAndSwapInt32(&val, 13, 15)
	fmt.Println(swapped)
}
