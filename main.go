package main

import (
	"fmt"
	"os"

	"github.com/k4droid3/AoC/y24"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("No arguments provided.")
		os.Exit(1)
	}

	if os.Args[1] == "24" {
		y24.Solution1()
		y24.Solution2()
		y24.Solution3()
	}
}
