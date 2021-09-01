package main

import (
	"fmt"
	"sync"
)

// exercise
func printer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("print example")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printer(&wg)
	go printer(&wg)
	wg.Wait()
}
