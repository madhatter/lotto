package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// random is a random number generator
var random *rand.Rand

// init initializes the random number generator
func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// createSlice creates a slice of integers from 1 to max
func createSlice(max int) []int {
	min := 1
	arr := make([]int, max)

	for i := range arr {
		arr[i] = min + i
	}
	return arr
}

// generateNumbers generates n unique random numbers from 1 to l
func generateNumbers(n int, l int) []int {
	numbrs := createSlice(l)
	rand.Shuffle(len(numbrs), func(i, j int) { numbrs[i], numbrs[j] = numbrs[j], numbrs[i] })

	return numbrs[0:n]
}

// sortNumbers sorts a slice of integers in ascending order
func sortNumbers(numbers []int) []int {
	// Bubble sort
	for i := range len(numbers) {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

// generateNumbersString generates a string of n unique random numbers from 1 to l
func generateNumbersString(n int, l int) string {
	numbrs := sortNumbers(generateNumbers(n, l))
	var txt string

	for i := range len(numbrs) {
		if i == 0 {
			txt = strconv.Itoa(numbrs[i])
		} else {
			txt = txt + " " + strconv.Itoa(numbrs[i])
		}
	}
	return txt
}

// the main function
func main() {
	arg, _ := strconv.Atoi(os.Args[1])

	log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.DebugLevel)
	//log.Debugf("Argument: %v", arg)

	fmt.Printf("Set of %v.\n", arg)
	for i := 1; i < arg+1; i++ {
		fmt.Println(generateNumbersString(5, 50) + " | " + generateNumbersString(2, 12))
	}
}
