package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	c, _ := ioutil.ReadFile("./../input.txt")
	grid := make([][]rune, 0)
	for i, line := range bytes.Split(c, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune{})
		for _, c := range []rune(string(line)) {
			grid[i] = append(grid[i], c)
		}
	}

	stepRow := 0
	stepCol := 0
	trees := 0
	for {
		stepCol += 3
		stepRow++

		if stepCol > len(grid[0])-1 { //overstep x
			stepCol = (stepCol - (len(grid[0])))
		}

		if stepRow > len(grid)-1 {
			break
		}
		if grid[stepRow][stepCol] == 35 {
			trees++
		}

	}

	fmt.Println("Total Trees :-> ", trees)

}
