/*
 *
 */
package main

// Imports
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Product definition
type Product struct {
	id int
}

func (p *Product) print(prodCon string, start, end time.Time) {
	fmt.Println("Produto", p.id, "processado por ", prodCon, " com sucesso.\nInício: ", start, "\nTérmino: ", end, "\n--------------------------------------------------------------")
}

/**
 * Consume channel items.
 */
func consumer(id string, prodch <-chan Product, wg *sync.WaitGroup) {
	var start time.Time
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
		start = time.Now().UTC()

		// Sleep for 500 milliseconds
		time.Sleep(500 * time.Millisecond)

		// Get the end time
		end = time.Now().UTC()

		// Prints the log
		prod.print(id, start, end)
	}

	// Decrement the waitgroup counter
	wg.Done()
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
	// Create a waitgroup
	var wg sync.WaitGroup
	// Initiate the number of consumers
	var consumers = 5
	// Initiate the number of products
	var products = 10
	// Create a product channel
	cs := make(chan Product)

	//inititate the waitgroup with the consumers number
	//when the count gets to 0 all goroutines blocked are released
	wg.Add(consumers)

	// Starts all consumers routines
	// Needs to be called before the product creation to avoit deadlock
	for i := 0; i < consumers; i++ {
		var id = "consumidor" + strconv.Itoa(i)
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

	fmt.Printf("Todos os produtos foram consumidos. \n")
}
