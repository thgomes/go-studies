package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}

}

func main() {
	data := make(chan int)
	go worker(1, data)

	for i := 0; i < 10000; i++ {
		go worker(i, data)
	}

	for i := 0; i < 50000; i++ {
		data <- i
	}
}
