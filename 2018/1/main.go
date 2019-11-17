package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	freq := 0
	for scanner.Scan() {
		t := scanner.Text()
		n, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		freq += n
	}
	fmt.Printf("%d\n", freq)
}
