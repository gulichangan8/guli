package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

func add(ch chan int64) {
	for i := 0; i <= 5000; i++ {
		_, ok := <-ch
		if ok == false {
			break
		}
		if i == 5000 {
			close(ch)
			break
		}
		x = x + 1
		ch <- x
	}
	wg.Done()
}
func main() {
	var ch = make(chan int64)
	wg.Add(2)
	go add(ch)
	go add(ch)
	ch <- x
	wg.Wait()
	fmt.Println(x)
}
