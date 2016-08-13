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
	repeat_max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Usage: %s <repeat count> <program path>\n", os.Args[0])
	}
	program := os.Args[2]
	for count := 0; count < repeat_max; count++ {
		cmd := exec.Command(program)
		var out bytes.Buffer
		cmd.Stdout = &out
		log.Println(time.Now())
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Damn, there was an error executing %s", program)
		}
		log.Println(time.Now())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out.String)
	}
}
