package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ResCal processing ...")
	vegeta := os.Args[1:]
	dockerStat := os.Args[2:]

	if len(vegeta) < 1 || len(dockerStat) < 1 {
		fmt.Println("You must pass two file result from vegeta[1] and docker stat[2]")
	} else {
		ds := readDockerStat(dockerStat[0])
		v := readVegeta(vegeta[0])

		fmt.Println(ds)
		fmt.Println(v)
	}
}
