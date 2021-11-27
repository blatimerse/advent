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

	m := make(map[byte]int)
	count := 0
	lines := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if len(s) != 0 {
			lines++
			for _, c := range []byte(s) {
				m[c]++
			}
			continue
		}
		count += addup(m, lines)
		lines = 0
		m = make(map[byte]int)
	}

	// if last line wasn't blank, need to add it, too..
	count += addup(m, lines)
	fmt.Println(count)
}

func addup(m map[byte]int, lines int) int {
	count := 0
	for _, v := range m {
		if v == lines {
			count++
		}
	}
	//fmt.Printf("%d of %d lines; m is %v\n", count, lines, m)
	return count
}
