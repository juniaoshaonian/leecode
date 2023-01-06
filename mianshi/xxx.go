package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	wg := sync.WaitGroup{}

	go func() {
		for {
			select {
			case num := <-ch1:
				printNum(num)
				wg.Done()
			}
		}
	}()
	go func() {
		for {
			select {
			case num := <-ch2:
				printNum(num)
				wg.Done()
			}
		}
	}()
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		if i%2 == 0 {
			ch2 <- i
		} else {
			ch1 <- i
		}
		fmt.Println(i)
	}
	wg.Wait()
}

func printNum(i int) {
	fmt.Println(i)
}
