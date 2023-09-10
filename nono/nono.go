package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

// Thread T1
func main() {
	channel := make(chan int)
	for i := 0; i < 1000; i++ {
		go worker(i, channel) // Thread Ti
	}

	for i := 0; i < 10000; i++ {
		channel <- i
	}
}
