package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	// consumer
	go func() {
		ticker := time.NewTicker(time.Second)

		for _ = range ticker.C {
			select {
			case <-done:
				println("child process interrupt...")
				return

			default:
				fmt.Printf("receive message: %d\n", <-messages)
			}
		}
	}()

	for i := 0; i < 20; i++ {
		messages <- i
		fmt.Printf("Putting %d\n", i)
		time.Sleep(time.Second / 2)
	}

	time.Sleep(2 * time.Second)
	close(done)
	time.Sleep(time.Second)
	println("main exit")
}


