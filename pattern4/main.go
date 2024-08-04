package main

import (
	"fmt"
	"sync"
)

func worker(done chan struct{}, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val := <-jobs:
			fmt.Println("Получено задание ", val)
		case <-done:
			fmt.Println("Получен сигнал отмены")
			return
		}
	}
}

func main() {
	jobs := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(done, jobs, &wg)
	for i := 0; i < 100; i++ {
		if i > 10 {
			close(done)
			break
		}
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
