package main

import (
	"fmt"
	"time"
)

type queue struct {
	pauses chan chan struct{}
}

func newQueue() *queue {
	return &queue{
		// Allocate the pause channel used to communicate the continuation channel.
		pauses: make(chan chan struct{}),
	}
}

func (q *queue) execute() {
	for {
		// The central part of the queue.
		// Listens on the pause channel for pause() to pass a channel through.
		// The loop then listens (breaks) on the channel waiting for it to be closed externally.
		select {
		case c := <-q.pauses:
			fmt.Println("Pausing")
			<-c // Block until the channel is closed.
			fmt.Println("Continuing")
		case <-time.After(time.Millisecond * 100):
			fmt.Println("Tick")
		}
	}
}

func (q *queue) pause() chan struct{} {
	// Allocate the continue channel used to continue the queue when time comes.
	c := make(chan struct{})

	// Send the created channel through the pause channel.
	q.pauses <- c

	// Return the continue channel back to the client.
	return c
}

func main() {
	// Create a new queue and start working.
	q := newQueue()
	go q.execute()

	//  Client loop interacting with the queue.
	for {
		// Pausing the queue returns channel that when closed continues
		// the paused queue. Failing to do so will cause the queue to
		// halt forever.
		<-time.After(time.Second * 2)
		c := q.pause()
		<-time.After(time.Second * 2)
		close(c)
	}
}
