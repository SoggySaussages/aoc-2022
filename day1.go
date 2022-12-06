package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const ElvesToSum = 3

func main() {
	input, err := os.Open("Inputs/day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	var currentCalories int
	var elves []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, currentCalories)
			currentCalories = 0
		} else {
			incr, _ := strconv.Atoi(scanner.Text())
			currentCalories += incr
		}
	}
	elves = append(elves, currentCalories)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println(sumSlice(elves[:ElvesToSum]))
}

func sumSlice(slice []int) (sum int) {
	for _, num := range slice {
		sum += num
	}
	return
}
