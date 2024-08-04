package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val := <-jobs:
			fmt.Println("Get value from channel ", val)
		case <-ctx.Done():
			fmt.Println("Context has been finished")
			return
		}
	}
}

func main() {
	jobs := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, jobs, &wg)
	go func() {
		for i := 0; i < 100; i++ {
			jobs <- i
			time.Sleep(time.Millisecond * 500)
		}
	}()

	wg.Wait()
}
