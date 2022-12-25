package main

import (
	"fmt"
	"sync"
)

var (
	ch  = make(chan int)
	ch2 = make(chan int)
	wg  sync.WaitGroup
)

func main() {
	wg.Add(1)
	go cal1()
	go cal2()
	wg.Wait()
}

func cal1() {
	ch <- 1
	for i := 0; i < 10; i++ {
		println(i)
	}
}
func cal2() {
	<-ch
	var i int64 = 10
	for ; i < 1000000; i++ {
		var sum, n int64 = 0, 2
		temp := i
		temp = temp / 10
		for ; n < 7; n++ {
			temp = temp / 10
			if temp < 1 {
				break
			}
		}
		for j := 0; j < 2; j++ {
			sum += (i % 10) * n
		}
		if sum == i {
			fmt.Println(i)
		}
	}

}
