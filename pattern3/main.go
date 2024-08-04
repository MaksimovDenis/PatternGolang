package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1, ch2 := generator(), generator()
OUT:
	for {
		select {
		case val, ok := <-ch1:
			fmt.Println("Канал 1", val)
			if !ok {
				break OUT
			}

		case val, ok := <-ch2:
			fmt.Println("Канал 2", val)
			if !ok {
				break OUT
			}
		}
	}

}

func generator() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- rand.Intn(100)
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()
	return ch
}
