package main

import (
	"fmt"
)

// Генератор выдаёт последовательность
// исходных данных.
func generator() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

// Процессор 1.
func processor1(src chan int) chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for n := range src {
			res <- n * n
		}
	}()
	return res
}

// Процессор 2.
func processor2(src chan int) chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for n := range src {
			res <- n * 2
		}
	}()
	return res
}

func main() {
	fmt.Println("Конвейер 1")
	ch := processor2(processor1(generator()))
	for n := range ch {
		fmt.Printf("получен результат работы конвейера: %d\n", n)
	}

	fmt.Println("Конвейер 2")
	// Процессоры можно как угодно комбинировать.
	ch = processor1(processor2(processor1(generator())))
	for n := range ch {
		fmt.Printf("получен результат работы конвейера: %d\n", n)
	}
}
