package main

import (
	"fmt"
	"time"
)

type Worker struct {
	id int
}

func main() {
	c := make(chan int, 100)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 10)
	}
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("Worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 10)
	}
}
