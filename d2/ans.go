package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// check verifies if a report is safe
func check(a []int) bool {
	// Determine the direction: increasing (1) or decreasing (-1)
	direction := a[1] - a[0]
	if direction == 0 {
		return false // No change in levels
	}
	direction /= int(math.Abs(float64(direction)))

	for i := 1; i < len(a); i++ {
		difference := a[i] - a[i-1]

		// Check if difference is within the allowed range
		if math.Abs(float64(difference)) < 1 || math.Abs(float64(difference)) > 3 {
			return false
		}

		// Check if the direction is consistent
		currentDirection := difference / int(math.Abs(float64(difference)))
		if currentDirection != direction {
			return false
		}
	}
	return true
}

func main() {
	var cnt int
	inp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	inpscan := bufio.NewScanner(inp)
	for inpscan.Scan() {
		line := inpscan.Text()
		fields := strings.Fields(line)

		// Parse the levels into an integer slice
		var a []int
		for _, v := range fields {
			level, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			a = append(a, level)
		}

		// Check if the report is safe
		if check(a) {
			cnt++
		}
	}

	if err := inpscan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cnt)
}
