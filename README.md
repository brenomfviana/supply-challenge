# Supply Challenge

## Description

This is a work from the Concurrent Programming course where the solution to the problem was implemented using the Go programming language.

## Problem

### Description

One of the biggest challenges in today's highly competitive marketplace is the _Supply Challenge_. Companies compete innovatively and aggressively so that their products and services are delivered to customers in the shortest time and at the lowest possible cost. In this scenario, the use of computational solutions and tools has shown to be a significant strategic differential. More often, companies are investing in modernizing their support systems to stand out from their market competitors.

One of the companies directly involved in this dispute is _ATL Transportes e Logística_. This transport company is responsible for delivering the products of a number of major suppliers and it is historically highly regarded in its market segment. However, with the advent of Internet sales, the company has realized that its order-processing capability is gradually becoming deprecated. If it continues declining at this rate, in a short time ATL will lose its prominence in the market achieved over the years.

Realizing the need to change its operations, ATL decided it was time to review the order fulfillment process. Currently the receipt of orders for delivery of products is made by phone and/or your web page. While allowing multiple requests to be made at the same time in a capacity of 50 telephone attendants and 300 simultaneous requests over the Internet, the internal processing of requests is done by a _software_ system that analyzes them one at a time. This was the constraint on the flow of operations identified by ATL's internal IT staff.

As all of the IT human resources of the company are allocated in other projects and have little or no time to carry out this modernization, the ATL's IT director decided to open a competition in the market so that several companies would present a solution to improve the capacity Order processing. The general idea is to implement an asynchronous concurrent solution for order processing following the general principles of the _Problem Producers and Consumers_, classic in competing systems.

For the competition, ATL does not require the implementation of a complete _software_ system, but only a prototype that makes it possible to evidence the results of a future formal implementation. The requests will be sent in a data format consisting of a numeric identifier and a data packet in text format. Customers will make use of the new tool to feed an internal _buffer_ with capacity for 5000 requests. Asynchronous internal processing will consume the requests and process them individually, so customers will not wait online for confirmation, but will receive a response later.

### Tasks

The central task to be accomplished in this work is to design and implement in the Go programming language the prototype required by ATL for order processing using concurrent programming concepts and techniques. The solution should be developed in three stages, each resulting in a version of the final solution implementation, which should be retained for benchmarking purposes.

The implementation should ensure that the program is competitive with the competition and properly apply the gorotine facilities. To promote synchronization between gorotines, you can use communication channels (with or without _buffer_) and/or advanced synchronization mechanisms provided by the language sync package. Your source code should be commented on in order to provide an adequate understanding of what was done.

#### Phase 1

Create gorotine-based agents that consume a 5000-position _buffer_ previously populated with requests (that is, only consumer agents should be implemented at this stage). The consumer will continually delete requests from the _buffer_ and its processing time should be simulated by a 500-millisecond pause. At the end of each processing, a _log_ should be displayed with the consumer agent identification, order identification, start time and processing end time. When the _buffer_ is empty, the consumer agents will be blocked. For now there is no need to worry about the necessary synchronization between the various consumer agents.

Once the implementation is complete, perform an experiment in which different quantities of concurrent consumer agents are started varying from 1, 5, 10, 50, 100, 500 and 1000. In order to obtain greater statistical significance, execute at least ten executions for each of these quantities and store the execution time of each of them. To display the results, create a table that lists the mean, minimum, and maximum times, plus the standard deviation of the ten runs for each quantity of consumer agents. Also produce a chart showing the average running times for each quantity of consumer agents.

#### Phase 2

Change the implementation developed in the previous phase so that the _buffer_ is now filled by a producer agent (also based on gorotines) so that there is synchronization between producer and consumer agents. Producing agents should create an order and their processing time should be simulated by a 500 millisecond pause. Each producer agent will feed the _buffer_ continuously as long as there is room in it and should be locked when full. When there is new available _buffer_ space resulting from the eventual consumption of an order by a consumer agent, the producer agent must be reactivated. At the end of each processing, a _log_ must be displayed with the producer or consumer agent ID, order identification, start time and processing time.

Once the implementation is complete, perform a new experiment in which different quantities of both producing and consuming agents are started, varying from 1, 5, 10, 50, 100, 500 and 1000. In order to obtain greater statistical significance, execute each of these quantities at least ten times and store the run time for each of them. To display the results, create a table that lists the average, minimum, and maximum times, plus the standard deviation of the 10 executions for each quantity of producer and consumer agents. Also draw a graph showing the average execution times for each number of agents.

**Tip**: The stopping criterion for the experiments depends significantly on what data structure will be used to represent the _buffer_ and on how the consumer agents perform their execution on it. There are at least two possible scenarios:

1. If consumers are literally removing items from the _buffer_, producers will hardly be able to fill them up to the 5000th position. In this case, there could be a sort of counter that, when it has reached 5000, producers will stop producing new items. Because consumers are removing items from the _buffer_, they are suspended when the _buffer_ is empty.

2. If the _buffer_ size is kept constant at 5000 (for example, if the _buffer_ has been implemented as a conventional _array_), producers should be suspended when the _buffer_ is full. In this implementation, the action of a consumer would mark the item of the buffer as consumed, so that consumers would be suspended when there are no more items to be consumed.

#### Phase 3

Change the implementation developed in the previous phase to meet a new requirement placed by the ATL IT directorate: we now want to ensure that each request is processed in the order it was requested in a FIFO policy (_first-in_, _first-out_). You should enhance your consumer agent implementation to ensure this specification.

After the implementation activity, retry the experiment performed in the previous step, now with the ordered processing guarantee. For discussion, compare the results now obtained with those obtained previously and analyze the impact (if any) presented by including this new one, not only in terms of performance, but also in terms of code simplification and future ease of maintenance.

## Dependencies

- Go 1.8.3

## Running

To run only one instance for a specific phase run the ```run.sh``` script.

```bash
# No need to enter the number of producers for the first step.
./run.sh <phase-id> <number-of-consumers> <number-of-producers>
```

To run all instances for a specific phase run the ```benchmarking.sh``` script.

```bash
./benchmarking.sh <phase-id>
```

## Team

[<img src="https://avatars2.githubusercontent.com/u/17532418?v=3&s=400" width="100"/>](https://github.com/brenomfviana) | [<img src="https://avatars2.githubusercontent.com/u/8813353?v=3&s=400" width="100"/>](https://github.com/Barbalho12) | [<img src="https://avatars2.githubusercontent.com/u/17392686?v=3&s=400" width="100"/>](https://github.com/Pekorishia)
---|---|---
[Breno Viana](https://github.com/brenomfviana) | [Felipe Barbalho](https://github.com/Barbalho12) | [Patrícia Cruz](https://github.com/Pekorishia)
