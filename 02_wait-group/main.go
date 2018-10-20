package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go fizz()
	go buzz()
	wg.Wait()
}

func fizz() {
	for i := 0; i < 5; i++ {
		fmt.Println("fizz", i)
	}
	wg.Done()
}

func buzz() {
	for i := 0; i < 5; i++ {
		fmt.Println("buzz", i)
	}
	wg.Done()
}
