package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Старт процессора №%d\n", id)
	for job := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("Процессор №%d выполнил задание №%d\n", id, job)
		result <- 0
	}
}

func main() {
	W := runtime.NumCPU()
	fmt.Println(W)
	const N = 10
	jobs := make(chan int)
	result := make(chan int)
	var wg sync.WaitGroup

	wg.Add(W)

	for i := 0; i < W; i++ {
		go worker(i, jobs, result, &wg)
	}

	go func(ch chan int) {
		for val := range ch {
			_ = val
		}
	}(result)

	for i := 0; i < N; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
