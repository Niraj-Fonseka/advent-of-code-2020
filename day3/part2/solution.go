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

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	multiTrees := 1
	for _, slope := range slopes {
		trees := traverse(grid, slope[0], slope[1])
		multiTrees *= trees
		fmt.Println("Total Trees :-> ", trees)
	}

	fmt.Println("final multiplied --> ", multiTrees)

}

func traverse(grid [][]rune, y, x int) int {
	stepRow := 0
	stepCol := 0
	trees := 0
	for {
		stepCol += y
		stepRow += x

		if stepCol > len(grid[0])-1 { //overstep x
			stepCol = (stepCol - (len(grid[0]))) //index reset
		}

		if stepRow > len(grid)-1 {
			break //reached / overstepping end
		}
		if grid[stepRow][stepCol] == 35 {
			trees++
		}
	}

	return trees
}
