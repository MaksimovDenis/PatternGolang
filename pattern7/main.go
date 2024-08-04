// Fan-in
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(nums []int) chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n * n
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()
	return ch
}

func fanIn(chans []chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(ch chan int) {
			defer wg.Done()
			for n := range ch {
				res <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	var chans []chan int

	for i := 0; i < 3; i++ {
		ch := worker([]int{1, 2, 3, 4, 5})
		chans = append(chans, ch)
	}

	ch := fanIn(chans)
	for n := range ch {
		fmt.Printf("Получено из результирующего канала: %d\n", n)
	}
}
