package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享的资源
var (
	sum   int
	mutex sync.RWMutex
)

func main() {
	// run()
	// doOnce()
	race()
}

func run() {
	var wg sync.WaitGroup
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}
	// 一直等待，只要计数器值为0
	wg.Wait()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	// 用于等待协程执行完毕
	done := make(chan bool)
	// 启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			// 把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

// 10个人赛跑，1个裁判发号施令
func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait()
			fmt.Println(num, "号开始跑...")
			cond.L.Unlock()
		}(i)
	}
	// 等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast()
	}()
	// 防止函数提前返回退出
	wg.Wait()
}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

// 增加了一个读取sum的函数，便于演示并发
func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}
