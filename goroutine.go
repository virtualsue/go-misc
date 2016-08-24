package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Start a goroutine and execute println concurrently
	for i := 0; i < 10; i++ {
		go message("goroutine " + strconv.Itoa(i))
	}
	fmt.Println("main function message")
	time.Sleep(10 * time.Second)
}

func message(msg string) {
	fmt.Println(msg)
}
