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

	fmt.Println(grid)
	// for row := 0; row < len(grid[0]); row++ {

	// 	for col := 0; col < len(grid); col++ {
	// 		fmt.Printf("%c ", grid[row][col])
	// 	}
	// 	fmt.Println()
	// }

	stepRow := 0
	stepCol := 0
	trees := 0
	for {
		stepCol += 3
		stepRow++

		// fmt.Println(len(grid[0])) -> 11
		if stepCol > len(grid[0])-1 { //overstep x
			//fmt.Println("Fixing....")
			//fmt.Printf(" %d - %d = %d \n", stepCol, (len(grid[0]) - 2), (stepCol - (len(grid[0]) - 2)))
			stepCol = (stepCol - (len(grid[0])))
		}

		//fmt.Printf("Col : %d , Row : %d  , Char : %c \n", stepCol, stepRow, grid[stepRow][stepCol])

		if stepRow > len(grid)-1 {
			break
		}
		//fmt.Printf("Checking X : %d , Y : %d  --> %c \n", stepCol, stepRow, grid[stepRow][stepCol])
		if grid[stepRow][stepCol] == 35 {
			fmt.Printf("Found a tree at col : %d , row : %d \n", stepCol, stepRow)
			trees++
		}

		//time.Sleep(1 * time.Second)
	}
	fmt.Println("Total trees --> ", trees)

	// y := 0
	// x := 0
	// trees := 0

	// for {

	// 	if y < len(grid) {
	// 		if grid[y][x%len(grid[0])] == '#' {

	// 			//fmt.Printf("Found a tree at col : %d , row : %d , char : %c \n", col%len(grid[0]), row, grid[row][col%len(grid[0])])
	// 			trees++
	// 		}

	// 		y += 1
	// 		x += 3
	// 	} else {
	// 		break
	// 	}

	// }

	fmt.Println("Total Trees :-> ", trees)

}
