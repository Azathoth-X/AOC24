package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var a []int
	var b []int
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	hash := make(map[int]int)
	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)
		// arr := strings.Split(line, " ")
		arr := strings.Fields(line)
		fmt.Println(arr)
		aa, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, aa)

		bb, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}
		hash[bb]++
		b = append(b, bb)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(a)
	slices.Sort(b)

	var sum int = 0
	var mult int = 0
	for i, _ := range a {
		sum += int(math.Abs(float64(a[i] - b[i])))

		mult += a[i] * hash[a[i]]
	}
	fmt.Println(sum)
	fmt.Println(mult)
}
