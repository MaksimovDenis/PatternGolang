// Fan-out
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(jobs chan int, num int, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	defer wg.Done()

	for job := range jobs {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		fmt.Printf("Поток %d отработал задание %d.\n", num, job)
	}
}

func main() {

	jobs := make(chan int)
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go worker(jobs, i, &wg)
	}
	for _, n := range []int{1, 2, 3, 4, 5} {
		jobs <- n
	}
	close(jobs)
	wg.Wait()
}
