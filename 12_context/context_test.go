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

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second)
		t.Log("准备调用cancel了")
		cancel()
	}()
	<-ctx.Done()
	t.Log("已经cancel了")
	t.Log(ctx.Err())
}

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	<-ctx.Done()
	t.Log("超时了")
	t.Log(ctx.Err())
}

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

func TestContextParentValue(t *testing.T) {
	parent := context.WithValue(context.Background(), "key1", "value1")
	son := context.WithValue(parent, "key1", "son-value1")
	t.Log(son.Value("key1"))
	_ = context.WithValue(parent, "key2", "son-value2")
	t.Log(parent.Value("key2"))
}
