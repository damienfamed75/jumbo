package main

// // #cgo LDFLAGS: -L./lib -lhello

/*
#include "./lib/hello.c"
#include "./lib/hello.h"
*/
import "C"

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	C.alphabet()

	fmt.Println("Finished!")
}

func calpha() {
	C.alphabet()
}

func alphabet() {
	input, err := os.OpenFile("~/input", os.O_RDONLY, 0666)
	buf := bufio.NewReader(input)

	out, err := os.OpenFile("~/output", os.O_CREATE|os.O_WRONLY, 0666)

	var (
		line []byte
	)

	for {
		line, err = buf.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("error whilst reading input file: %v", err)
		}

		out.Write(line)
	}

	// Attempt to clear memory early.
	line = nil

	// Not deferring because of performance issues.
	input.Close()
	out.Close()
}
