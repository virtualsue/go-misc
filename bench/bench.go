package main

// This program does simple execution timing of a program specified
// on the command line, repeated as many times as specified.
// It isn't a real benchmarking tool, it's for demo purposes.
//
// Usage: bench <# iterations> </some/program/to/time> <arg1> ... <argN>

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		Usage()
	}
	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Usage()
	}
	program := os.Args[2]
	opts := os.Args[3:]
	for i := 1; i <= iterations; i++ {
		start := time.Now()
		output, err := exec.Command(program, opts...).Output()
		duration := time.Since(start)
		fmt.Print(string(output))
		if err != nil {
			log.Fatalf("Damn, there was an error executing %s: %s", program, err)
		}
		fmt.Println(i, duration)
	}
}

func Usage() {
	log.Fatalf("Usage: %s <# iterations> <program path> <arg1> ... <argN>\n", os.Args[0])
}
