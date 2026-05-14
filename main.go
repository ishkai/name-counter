package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: name-counter <file_name> [--sort]")
		return
	}

	fileName := os.Args[1]
	sortByFrequency := false

	if len(os.Args) >= 3 {
		switch os.Args[2] {
		case "--sort":
			sortByFrequency = true
		default:
			fmt.Println("Unknown argument:", os.Args[2])
			return
		}
	}

	counts, err := CountNames(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	PrintResult(counts, sortByFrequency, os.Stdout)
}
