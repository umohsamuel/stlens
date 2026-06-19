package main

import (
	"fmt"
)

func main() {
	nurl, err := normalizeURL("https://www.boot.dev/blog/path")

	if err != nil {
		fmt.Printf(`error in %v`, err.Error())
	}

	fmt.Printf(`response is %v`, nurl)
}
