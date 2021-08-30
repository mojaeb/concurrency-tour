package main

import (
	"fmt"
)

func sender(message chan<- string) {
	message <- "hello"
}

func main() {
	message := make(chan string)

	go sender(message)
	go sender(message)

	fmt.Println(<-message)
	fmt.Println(<-message)
}
