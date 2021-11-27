package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

	seats := loadSeatData(in)

	max := -1
	for _, s := range seats {
		if s > max {
			max = s
		}
	}

	sort.IntSlice(seats).Sort()

	fmt.Println("max seat number: ", seats[len(seats)-1])

	prev := -1
	for _, i := range seats {
		if i != prev+1 && prev != -1 {
			fmt.Println("my seat is ", i-1)
			break
		}
		prev = i
	}
}

func loadSeatData(in io.Reader) []int {
	scanner := bufio.NewScanner(in)
	var seats []int
	for scanner.Scan() {
		s := scanner.Text()

		seat := findSeat(s)
		seats = append(seats, seat)
	}

	return seats
}

func findRow(s string) int {
	row := 0
	for _, x := range s {
		switch x {
		case 'F': // front half
			row <<= 1
		case 'B': // back half
			row <<= 1
			row++
		default:
			continue
		}
	}
	return row
}

func findColumn(s string) int {
	col := 0
	for _, x := range s {
		switch x {
		case 'L': // front half
			col <<= 1
		case 'R': // back half
			col <<= 1
			col++
		default:
			continue
		}
	}
	return col
}

func findSeat(s string) int {
	return 8*findRow(s) + findColumn(s)
}
