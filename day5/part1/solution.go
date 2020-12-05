package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

func main() {

	c, _ := ioutil.ReadFile("./../input.txt")
	maxSeat := 0
	for _, line := range bytes.Split(c, []byte("\n")) {
		if len(line) == 0 || len(line) > 10 {
			continue
		}

		rows := line[:7]
		cols := line[7:]

		midpoint := 127
		midpointSides := 7

		row := traverse(rows, 'F', 'B', 0, float64(0), float64(127), float64(midpoint))
		col := traverse(cols, 'L', 'R', 0, float64(0), float64(7), float64(midpointSides))

		seatID := int((row * 8) + col)

		fmt.Printf("Boarding Pass : %s , Seat ID : %d \n", line, seatID)
		if maxSeat < seatID {
			maxSeat = seatID
		}
	}

	fmt.Println("\nMax Seat : ", maxSeat)
}

func traverse(vals []byte, letterA byte, letterB byte, index int, min float64, max float64, midpoint float64) float64 {

	midpoint = math.Floor((max - min) / 2)

	if vals[index] == letterA {
		max = (max - midpoint) - 1
	}

	if vals[index] == letterB {
		min = (min + midpoint) + 1
	}

	if index == len(vals)-1 { //basecase
		return math.Min(min, max)
	}

	index++
	return traverse(vals, letterA, letterB, index, min, max, midpoint) //recurse
}
