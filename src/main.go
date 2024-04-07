package main

import (
	"fmt"
	"time"
)

func main() {
	ch := createIntChan(1000)
	fmt.Println("Hello, World!!!")
	go func(ch chan int) {
		for i := 0; i < 100; i++ {
			if i == 99 {
				i = 1
			}
			ch <- i
			time.Sleep(10 * time.Millisecond)
		}
	}(ch)
	go func(ch chan int) {
		for {
			select {
			case recv := <-ch:
				fmt.Printf("Rec: %d\n", recv)
			default:
				fmt.Println("Nothing")
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}(ch)

	time.Sleep(10000 * time.Second)
}

func createIntChan(n int) chan int {
	toReturned := make(chan int, n)
	return toReturned
}
