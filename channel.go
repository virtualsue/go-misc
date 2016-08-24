package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum	// send sum to channel
}

func main() {
	s := []int{1, 3, 5, 7, 9, 11}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //
	fmt.Println(x, y)
}
