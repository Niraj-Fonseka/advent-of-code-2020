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

	prepedData := dataPrep()

	bagType := "shiny_gold"
	var canCarry []string
	collected := make(map[string]bool)

	for bag, contains := range prepedData {
		_, ok := contains[bagType]
		if ok {
			collected[bag] = true
			canCarry = append(canCarry, bag)
		}
	}

	for {
		var collect []string
		for _, b := range canCarry {

			for bag, contains := range prepedData {
				_, ok := contains[b]
				if ok {
					collected[bag] = true
					collect = append(collect, bag)
				}
			}
		}

		if len(collect) == 0 {
			break
		}

		canCarry = make([]string, 0)
		copy(canCarry, collect)

		collect = []string{}

	}
	fmt.Println(len(collected))
}

func dataPrep() map[string]map[string]int64 {
	c, err := ioutil.ReadFile("./../input.txt")

	if err != nil {
		log.Fatal(err)
	}
	//maxSeat := 0
	bagStore := make(map[string]map[string]int64)
	for _, line := range bytes.Split(c, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		lineSTR := string(line)

		//fmt.Println(lineSTR)
		bag := strings.Replace(strings.TrimSpace(strings.Split(lineSTR, "bags contain")[0]), " ", "_", -1)
		contains := strings.TrimSpace(strings.Replace(strings.Replace(strings.TrimSpace(strings.Split(lineSTR, "bags contain")[1]), "bags", "", -1), ".", "", -1))

		_, ok := bagStore[bag]

		if !ok {
			bagStore[bag] = make(map[string]int64)
		}

		if contains == "no other" {
			continue
		}
		for _, contain_bag := range strings.Split(contains, ",") {
			bag_and_count := strings.Split(strings.TrimSpace(contain_bag), " ")
			bagCount, err := strconv.ParseInt(bag_and_count[0], 10, 0)

			bagType := fmt.Sprintf("%s_%s", bag_and_count[1], bag_and_count[2])

			if err != nil {
				log.Fatal(err)
			}

			bagStore[bag][bagType] = bagCount
		}
	}

	return bagStore
}
