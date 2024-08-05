package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("aoc.txt")
	check(err)

	var santaFloor int

	for _, r := range dat {
		c := string(r)
		if c == "(" {
			santaFloor++
		} else if c == ")" {
			santaFloor--
		} else {
			fmt.Println("Uncontrolled value:" + c)
		}
	}

	fmt.Println(santaFloor)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
