package main

import (
	"fmt"
	"sync"
)

func a() {
	var (
		ch = make(chan int)
		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		ch <- 1
		fmt.Println("出现")
		wg.Done()
	}()
	<-ch
	wg.Wait()
}
