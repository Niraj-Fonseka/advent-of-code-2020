package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	instructions []string
	values       []int
)

func main() {
	c, err := ioutil.ReadFile("./../input1.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range bytes.Split(c, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		instructionLine := strings.Split(string(line), " ")
		instructions = append(instructions, instructionLine[0])

		instructionCount, err := strconv.Atoi(instructionLine[1])
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, instructionCount)
	}

	visited := make([]int, len(values))

	accumulator := 0
	sp := 0

	for {

		fmt.Printf("processing instruction : %s , sp : %d , acc : %d \n", instructions[sp], sp, accumulator)

		if visited[sp] == 1 {
			break
		}
		visited[sp] = 1
		sp, accumulator = processInstruction(instructions[sp], sp, accumulator)

	}

	fmt.Println("Final Accumulator : ", accumulator)

}

func processInstruction(instruction string, sp int, acc int) (int, int) {

	switch instruction {
	case "nop":
		return sp + 1, acc
	case "jmp":
		return sp + values[sp], acc
	case "acc":
		return sp + 1, acc + values[sp]
	}

	return -99, -99
}
