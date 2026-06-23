package main

import (
	"fmt"
	"os"
)

// https://learnwebscraping.dev/practice/ecommerce/

// https://learnwebscraping.dev/practice/ecommerce/products/ashenfang-longsword-fan-1001/

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v", args[0])
}
