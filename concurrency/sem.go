package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		c <- i
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("ID %d started\n", i)
	time.Sleep(5 * time.Second)
	fmt.Printf("ID %d finished\n", i)
	<-c
}
