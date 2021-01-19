package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	ctx, stop := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, "userID", 2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗1】")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗2】")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗3】")
	}()
	time.Sleep(3 * time.Second)
	stop()
	wg.Wait()
}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】", "协程退出")
			return
		default:
			userID := ctx.Value("userID")
			fmt.Println("【获取用户】", "用户ID为:", userID)
			time.Sleep(1 * time.Second)
		}
	}
}

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}
