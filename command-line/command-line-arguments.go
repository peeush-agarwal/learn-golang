package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Printf("Args with Prog %v\n", argsWithProg)
	fmt.Printf("Args without Prog %v\n", argsWithoutProg)
	fmt.Printf("Args %v\n", arg)
}
