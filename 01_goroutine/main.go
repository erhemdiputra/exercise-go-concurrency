package main

import (
	"fmt"
	"time"
)

func main() {
	go fizz()
	go buzz()

	time.Sleep(time.Second)
}

func fizz() {
	for i := 0; i < 5; i++ {
		fmt.Println("fizz", i)
	}
}

func buzz() {
	for i := 0; i < 5; i++ {
		fmt.Println("buzz", i)
	}
}
