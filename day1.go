package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count1, count2 := 0, 0
	var vals [3]int64

	scanner.Scan()

	current, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	vals[0], _ = strconv.ParseInt(scanner.Text(), 10, 0)
	scanner.Scan()
	vals[1], _ = strconv.ParseInt(scanner.Text(), 10, 0)
	if vals[1] > vals[0] {
		count1++
	}
	scanner.Scan()
	vals[2], _ = strconv.ParseInt(scanner.Text(), 10, 0)
	if vals[2] > vals[1] {
		count1++
	}

	for scanner.Scan() {
		number, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		if number > current {
			count1 += 1
		}
		if number > vals[0] {
			count2 += 1
		}
		current = number
		vals[0] = vals[1]
		vals[1] = vals[2]
		vals[2] = number
	}

	message := fmt.Sprintf("%v single increases", count1)
	fmt.Println(message)
	message = fmt.Sprintf("%v sum increases", count2)
	fmt.Println(message)
}
