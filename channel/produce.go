package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Producer: ", i)
		queue <- i
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("receive:", v)
	}
}

func main() {
	queue := make(chan int, 1)
	go Producer(queue)
	go Consumer(queue)
	time.Sleep(1 * time.Second) //让Producer与Consumer完成
}
