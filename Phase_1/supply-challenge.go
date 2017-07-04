/*
 * Phase 1.
 */
package main

// Imports
import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

/**
 * Product definition
 */
type Product struct {
	id int
}

/**
 * Prints the log with the product ID, start time of production and end time of
 * production.
 */
func (p *Product) printLog(prodCon string, start, end time.Time) {
	fmt.Println("Product", p.id, "processed by ", prodCon, " successfully.", "\n",
		"Time Init: ", fmt.Sprintf("%dH%dm%ds", start.Hour(), start.Minute(), start.Second()), "\n",
		"Time End:  ", fmt.Sprintf("%dH%dm%ds", end.Hour(), end.Minute(), end.Second()), "\n",
		"--------------------------------------------------------------")
}

/**
 * Consume channel items.
 */
func consumer(id string, prodch <-chan Product, wg *sync.WaitGroup) {
	// Start time
	var start time.Time
	// End time
	var end time.Time
	// Infinite loop
	for {
		// Get the product and a boolean that verifies if the channel is open
		prod, open := <-prodch
		// If the channel is closed
		if !open {
			// Leave the loop
			break
		}
		// Get the start time
		start = time.Now()
		// Sleep for 500 milliseconds
		time.Sleep(500 * time.Millisecond)
		// Get the end time
		end = time.Now()
		// Prints the log
		prod.printLog(id, start, end)
	}
	// Decrement the waitgroup counter
	wg.Done()
}

/**
 * Start program execution.
 */
func main() {
	// Create a waitgroup
	var wg sync.WaitGroup
	// Initiate the number of consumers
	var consumers = 5
	// Initiate the number of products
	var products = 5000
	// Reading of arguments for number of consumers and producers, respectively
	if len(os.Args) >= 2 {
		consumers, _ = strconv.Atoi(os.Args[1])
	}
	// Create a product channel
	cs := make(chan Product)
	// Inititate the waitgroup with the consumers number
	// when the count gets to 0 all goroutines blocked are released
	wg.Add(consumers)
	// Starts all consumers routines
	// Needs to be called before the product creation to avoit deadlock
	for i := 0; i < consumers; i++ {
		var id = "consumidor " + strconv.Itoa(i)
		go consumer(id, cs, &wg)
	}
	// Creates all products
	for i := 0; i < products; i++ {
		var p Product
		p.id = i
		// Put the product inside the channel
		cs <- p
	}
	// Close the channel
	close(cs)
	// Make the main wait until the wg counter gets to zero
	wg.Wait()
	// End message
	fmt.Printf("Todos os produtos foram consumidos. \n")
}
