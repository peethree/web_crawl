package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

// usage: ./crawler URL maxConcurrency maxPages
func main() {

	// initialize a config struct object
	cfg := &config{
		pages: make(map[string]int),
		mu:    &sync.Mutex{},
		wg:    &sync.WaitGroup{},
	}

	// handle improper use
	cliArguments := os.Args[1:]

	if len(cliArguments) < 3 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	if len(cliArguments) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// get the given website url
	givenURL := cliArguments[0]

	//set maxConcurrency amount
	amount, err := strconv.Atoi(cliArguments[1])
	if err != nil {
		fmt.Println("Error converting second command line argument to int:", err)
	}
	cfg.concurrencyControl = make(chan struct{}, amount)

	//set maxPages
	max, err := strconv.Atoi(cliArguments[2])
	if err != nil {
		fmt.Println("Error converting third command line argument to int:", err)
	}
	cfg.maxPages = max

	// parse the given url into type url.URL
	parsedURL, err := url.Parse(givenURL)
	if err != nil {
		fmt.Println("error parsing the url given in the command line")
	}
	// set the cfg struct's baseURL
	cfg.baseURL = parsedURL

	fmt.Println("starting crawl of:", givenURL)
	// wait group before the go routine
	cfg.wg.Add(1)
	// go routine of recursive function
	go cfg.crawlPage(givenURL)
	// wait group wait after function
	cfg.wg.Wait()

	// for rcuNormalized, count := range cfg.pages {
	// 	fmt.Printf("%d - %s\n", count, rcuNormalized)
	// }

	printReport(cfg.pages, givenURL)
}
