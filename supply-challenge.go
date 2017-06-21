/*
 *
 */
package main

// Imports
import (
    "fmt"
    "time"
)

/**
 * Consume channel items.
 */
func consumer(ch chan int) {
	for i := range ch {
	    r := i
	    fmt.Println("Consumed: ", r)
	}
}

/**
 * Produce channel items.
 */
func producer(ch chan int) {
    //
}

/**
 * Start program execution.
 */
func main() {
    // Create a channel with a 5000 capacity buffer
    var ch chan int
    ch = make(chan int, 5000)
    // Add items to the channel
    for i := 0; i <= 10; i++ {
        ch <- i
	    fmt.Println("Produced: ", i)
    }
    // Close channel
    close(ch)
    // Start consumer
    go consumer(ch)
	time.Sleep(1e9)
}
