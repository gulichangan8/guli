package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func func1(A chan bool, B chan bool) {
	for i := 1; i <= 100; i += 2 {
		if flag := <-A; flag == true {
			fmt.Println(i)
			B <- true
		}
	}
	w.Done()
}

func func2(A chan bool, B chan bool) {
	for i := 2; i <= 100; i += 2 {
		if flag := <-B; flag == true {
			fmt.Println(i)
			if i != 100 {
				A <- true
			}
		}
	}
	w.Done()
}

func main() {
	A := make(chan bool)
	B := make(chan bool)
	w.Add(2)
	go func1(A, B)
	go func2(A, B)
	A <- true
	w.Wait()
}
