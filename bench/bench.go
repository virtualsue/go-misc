package main

import (
	"bytes"
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
	repeat_max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Usage()
	}
	program := os.Args[2]
	for count := 0; count < repeat_max; count++ {
		cmd := exec.Command(program)
		var out bytes.Buffer
		cmd.Stdout = &out
		log.Println(time.Now())
		err = cmd.Run()
		log.Println(time.Now())
		if err != nil {
			log.Fatalf("Damn, there was an error executing %s: %s", program, err)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out.String)
	}
}

func Usage() {
		log.Fatalf("Usage: %s <# iterations> <program path>\n", os.Args[0])
}
