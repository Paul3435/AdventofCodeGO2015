package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("aoc.txt")
	check(err)

	var santaFloor int
	var firstNegativeFloor int
	firstTimeNegative := false

	for i, r := range dat {
		c := string(r)
		if c == "(" {
			santaFloor++
		} else if c == ")" {
			santaFloor--
			if santaFloor == -1 && firstTimeNegative == false {
				firstNegativeFloor = i
				firstTimeNegative = true
			}
		} else {
			fmt.Println("Uncontrolled value:" + c)
		}

	}

	fmt.Println(santaFloor)
	fmt.Println(firstNegativeFloor + 1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
