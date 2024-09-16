package main

import (
	"fmt"
	"os"
)

func main() {
	cliArguments := os.Args[1:]

	if len(cliArguments) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(cliArguments) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := cliArguments[0]

	if len(cliArguments) == 1 {
		fmt.Println("starting crawl of:", baseURL)
		htmlBody, err := getHTML(baseURL)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(htmlBody)
	}
}
