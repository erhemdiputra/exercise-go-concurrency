package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("Num CPU = ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go fizz()
	go buzz()
	wg.Wait()
}

func fizz() {
	for i := 0; i < 5; i++ {
		fmt.Println("fizz", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func buzz() {
	for i := 0; i < 5; i++ {
		fmt.Println("buzz", i)
		time.Sleep(10 * time.Millisecond)
	}
	wg.Done()
}
