package _2_context

import (
	"context"
	"testing"
	"time"
)

func TestContextValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key1", "value1")
	val, ok := ctx.Value("key1").(string)
	t.Log(val, ok)
}

// 主动取消
func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second)
		t.Log("准备调用cancel了")
		// 关掉ctx.Done()的channel，从而进入到23行
		cancel()
	}()
	<-ctx.Done()
	t.Log("已经cancel了")
	t.Log(ctx.Err())
}

// 超时
func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	<-ctx.Done()
	t.Log("超时了")
	t.Log(ctx.Err())
}

// 控制是从上到下，父结束，子也结束
func TestContextParentCancel(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Second, func() {
		cancel()
	})
	son, sonCancel := context.WithCancel(parent)
	<-son.Done()
	t.Log("son已经过来了")
	sonCancel()
}

// panic
func TestContextParentCancel1(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	_, sonCancel := context.WithCancel(parent)
	time.AfterFunc(time.Second, func() {
		sonCancel()
	})
	<-parent.Done()
	t.Log("parent已经过来了")
	cancel()
}

// 取值是从下至上，儿子现在自己哪里找，找不到再去父亲哪里找
func TestContextParentValue(t *testing.T) {
	parent := context.WithValue(context.Background(), "key1", "value1")
	son := context.WithValue(parent, "key1", "son-value1")
	t.Log(son.Value("key1"))
	// key2在son中，parent取不到
	_ = context.WithValue(parent, "key2", "son-value2")
	t.Log(parent.Value("key2"))
}
