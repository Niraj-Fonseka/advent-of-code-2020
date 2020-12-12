package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	c, err := ioutil.ReadFile("./../input1.txt")

	if err != nil {
		log.Fatal(err)
	}
	//maxSeat := 0
	var instructions []string
	var values []int
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

	fmt.Println("Instructions : ", instructions)
	fmt.Println("Values : ", values)
	fmt.Println("Visited : ", visited)

}
