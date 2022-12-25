package main

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan int)
)

func b() {
	wg.Add(3)
	go Work("goroutine1")
	<-ch1
	go Work("goroutine2")
	<-ch1
	go Work("goroutine3")
	<-ch1
	wg.Wait()
	fmt.Println("successful")
}

func Work(workName string) {
	ch1 <- 1
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	wg.Done()
}
