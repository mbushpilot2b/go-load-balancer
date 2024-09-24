package main

import (
	"go-load-balancer/balancer"
	"math/rand"
	"time"
)

const nWorker = 3

func main() {
	work := make(chan balancer.Request)
	done := make(chan *balancer.Worker)

	// Create workers and start them
	workers := make([]*balancer.Worker, nWorker)
	for i := 0; i < nWorker; i++ {
		workers[i] = balancer.NewWorker(i, done)
		go workers[i].Work(done)
	}

	// Create balancer and start it
	b := balancer.NewBalancer(workers, done)
	go b.Balance(work)

	// Start requester
	go requester(work)

	// Run indefinitely
	select {}
}

func requester(work chan<- balancer.Request) {
	c := make(chan int)
	for {
		time.Sleep(time.Duration(rand.Int63n(nWorker * 2 * int64(time.Second))))
		work <- balancer.Request{Fn: workFn, C: c} // send request
		result := <-c                              // wait for answer
		furtherProcess(result)
	}
}

func workFn() int {
	// Simulate some work
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func furtherProcess(result int) {
	// Process the result
	println("Result:", result)
}
