package main

import (
	"fmt"
	"sync"
)

func do(jobs chan int, errors chan error) chan int {
	res := make(chan int)
	go func() {
		for n := range jobs {
			if n%5 == 0 {
				errors <- fmt.Errorf("%d don't devide on 5", n)

				continue
			}
			res <- n * n
		}
		close(res)
		close(errors)
	}()
	return res
}

func main() {
	jobs := make(chan int)
	errs := make(chan error)
	res := do(jobs, errs)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for n := range res {
			fmt.Println("Результат: ", n)
		}
	}()

	go func() {
		defer wg.Done()
		for err := range errs {
			fmt.Println("Ошибка: ", err)
		}
	}()

	for i := 0; i < 12; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

}
