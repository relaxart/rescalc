package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("ResCal processing ...")

	c := ParseConfig("./config.yaml")
	fmt.Printf("Resource deviation: %f%", c.Deviation)
	result := os.Args[1:]
	if len(result) < 1 {
		fmt.Println("You must pass results")
	} else {
		r := readResult(result[0])
		p := Process(c, r.Results)

		fmt.Println("Calculated resources:", "\n")
		ProcessResults(p)
	}
}
