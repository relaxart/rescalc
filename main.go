package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Add stress test file: ")
	vegetta, _ := reader.ReadString('\n')
	fmt.Println(vegetta)
	fmt.Print("Add config file: ")
	config, _ := reader.ReadString('\n')
	fmt.Println(config)
}
