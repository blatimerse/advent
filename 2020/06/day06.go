package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var in io.Reader = os.Stdin

	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		in = f
	}

	loadQuestionData(in)
}

func loadQuestionData(in io.Reader) {
	scanner := bufio.NewScanner(in)

	m := make(map[byte]struct{})
	count := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if len(s) != 0 {
			for _, c := range []byte(s) {
				m[c] = struct{}{}
			}
			continue
		}
		count += len(m)
		m = make(map[byte]struct{})
	}

	count += len(m)
	fmt.Println(count)
}
