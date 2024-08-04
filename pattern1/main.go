package main

import "fmt"

func Async(nums []int) chan int {
	ch := make(chan int)

	go func() {
		for _, n := range nums {
			ch <- n * n
		}
		close(ch)
	}()

	return ch
}

func main() {
	src := []int{1, 2, 3}

	res := []int{}
	ch := Async(src)
	for n := range ch {
		res = append(res, n)
	}
	fmt.Println(res)
}
