package main

// Imports
import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Global constant
const MAX_VALUE int = 5000

/**
 * Product definition
 */
type Product struct {
	id float32
}

/**
 * Counter definition
 */
type Counter struct {
	val    int
	finish bool
	closed bool
}

/**
 * Fifo definition
 */
type FIFO struct {
	muxPut    sync.Mutex
	muxRemove sync.Mutex
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
 * Put the product inside the channel and print the log.
 */
func (f *FIFO) putProduct(pId string, start time.Time, prod Product, prodch chan Product, c *Counter) bool {
	f.muxPut.Lock()
	// If the number of products is smaller than the max value
	if c.val < MAX_VALUE {
		c.val++
		c.closed = false
		c.finish = false
		// Get product
		prodch <- prod
		end := time.Now().UTC()
		// Print log
		prod.print(pId, start, end)
	} else {
		// The counter has finished
		c.finish = true
		// Check if another productor hasn't close the channel yet
		if !c.closed {
			c.closed = true
			// Close the channel
			close(prodch)
		}
	}
	defer f.muxPut.Unlock()
	return c.finish
}

/**
 * Remove the product of the channel and return the boolean.
 */
func (f *FIFO) removeProduct(cId string, start time.Time, prodch <-chan Product) bool {
	f.muxRemove.Lock()
	// Get the product and a boolean that verifies if the channel is open
	prod, open := <-prodch
	// If the channel is open
	if open {
		end := time.Now().UTC()
		prod.print(cId, start, end)
	}
	defer f.muxRemove.Unlock()
	return open
}

/**
 * Consume channel items.
 */
func consumer(id string, prodch <-chan Product, wg *sync.WaitGroup, fifo *FIFO) {
	var start time.Time
	// Infinite loop
	for {
		start = time.Now().UTC()
		// Wait
		time.Sleep(50 * time.Millisecond)
		// Get the boolean that verifies if the channel is open
		open := fifo.removeProduct(id, start, prodch)
		// If the channel is closed
		if !open {
			// Leave the loop
			break
		}
	}
	// Decrement the waitgroup counter
	wg.Done()
}

/**
 * Produce channel items.
 */
func producer(id int, prodch chan Product, wg *sync.WaitGroup, counter *Counter, fifo *FIFO) {
	var start time.Time
	count := 1
	// Infinite loop
	for {
		start = time.Now().UTC()
		// Wait
		time.Sleep(50 * time.Millisecond)
		// Create the product
		var prod Product
		// New product
		prod.id = float32(id) + (float32(count) / 1000)
		finish := fifo.putProduct("produtor "+strconv.Itoa(id), start, prod, prodch, counter)
		// If the counter finished
		if finish {
			break
		}
		count++
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
	// Create a counter
	var count Counter
	// Create a FIFO garanteer
	var fifo FIFO
	// Initiate the number of consumers
	var consumers = 5
	// Initiate the number of producers
	var producers = 5
	// Reading of arguments for number of consumers and producers, respectively
	if len(os.Args) >= 2 {
		consumers, _ = strconv.Atoi(os.Args[1])
	}
	if len(os.Args) >= 3 {
		producers, _ = strconv.Atoi(os.Args[2])
	}
	// Create a product channel
	cs := make(chan Product, 5000)
	// Inititate the waitgroup with the goroutines number
	// when the count gets to 0 all goroutines blocked are released
	wg.Add(consumers + producers)
	// Starts all consumers routines
	// Needs to be called before the product creation to avoid deadlock
	for i := 0; i < consumers; i++ {
		var id = "consumidor " + strconv.Itoa(i)
		go consumer(id, cs, &wg, &fifo)
	}
	// Create the producers
	for i := 0; i < producers; i++ {
		go producer(i, cs, &wg, &count, &fifo)
	}
	// Make the main wait until the wg counter gets to zero
	wg.Wait()
	fmt.Printf("Todos os produtos foram consumidos. \n")
}
