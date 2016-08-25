package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 30; x++ {
			naturals <- x
			time.Sleep(100 * time.Millisecond)
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer
	for {
		x, ok := <-squares
		if !ok {
			log.Println("Closing the program")
			return
		}
		fmt.Println(x)
	}
}
