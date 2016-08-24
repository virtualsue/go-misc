package main

import (
	"fmt"
	"os/exec"
)

func main() {
	n := 5000000
	delta := 1.0 / float64(n)
	var sum, x float64

	for i := 1; i <= n; i++ {
		x = (float64(i) - 0.5) * delta
		sum += 1.0 / (1.0 + x*x)
	}
	var pi = 4.0 * delta * sum
	fmt.Println("Pi is approximately", pi, "after", n, "iterations.")
}
