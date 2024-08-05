package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalPaper int

	for scanner.Scan() { //Will work as long as lines are shorted than 64k bytes, splits based on newLines
		stringSlice := strings.Split(scanner.Text(), "x") //make a slice using the x multiplication sign as separator
		intSlice := make([]int, len(stringSlice))
		for i, s := range stringSlice { //convert all strings to ints
			intSlice[i], _ = strconv.Atoi(s)
		}

		sort.Slice(intSlice, func(i, j int) bool { //sort slice by size
			return intSlice[i] < intSlice[j]
		})

		l := intSlice[0] //will always be the smaller
		w := intSlice[1] //always the second smallest
		h := intSlice[2]

		//totalPaper += 3*l*w + 2*w*h + 2*h*l
		totalPaper += 2*l + 2*w + l*w*h
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalPaper)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
