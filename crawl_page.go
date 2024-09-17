package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// decrement  wg counter by one.
	defer cfg.wg.Done()

	// rawcurrentURL has to be on same domain as rawBaseURL
	rcu, err := url.ParseRequestURI(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing raw url into a URL struct")
		return
	}

	// if the current url is not of the same domain as the baseURL
	if rcu.Host != "" && cfg.baseURL.Host != rcu.Host {
		fmt.Println("current url is not on the same domain")
		return
	}

	// get normalized version of rawcurrenturl
	rcuNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error with normalize raw current url")
	}

	// if the current url has already been visited: (isFirst = false) return
	if !cfg.addPageVisit(rcuNormalized) {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	// get the HTML from current url
	currentHTML, err := getHTML(rawCurrentURL)
	// fmt.Println("getting HTML from", rawCurrentURL)
	if err != nil {
		fmt.Println("error getting HTML:", err)
		return
	}

	// get the urls from currentHTML
	urls, err := getURLsFromHTML(currentHTML, cfg.baseURL.String())
	// fmt.Println("getting urls from html", urls)
	if err != nil {
		fmt.Println("error getting urls:", err)
		return
	}

	// crawl over every url in urls
	for _, u := range urls {
		cfg.wg.Add(1)
		// new go routine, anonymous function, calls itself at the end with (u)
		go func(url string) {
			// Block if too many active
			cfg.concurrencyControl <- struct{}{}
			// Release when done
			defer func() { <-cfg.concurrencyControl }()
			// recursive call
			cfg.crawlPage(url)
		}(u)
	}
}
