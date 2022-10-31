package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup

func add(c chan int) {
	for i := 0; i < 50000; i++ {
		c <- x
		x = x + 1
		<-c
	}
	wg.Done()
}
func main() {
	c := make(chan int, 1)
	wg.Add(2)
	go add(c)
	go add(c)
	wg.Wait()
	fmt.Println(x)
}
