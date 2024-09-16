package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// rawcurrentURL has to be on same domain as rawBaseURL
	rcu, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	rbu, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}

	// if the two given url's do not have the same domain name
	if rcu.Host != rbu.Host {
		fmt.Println("two given urls aren't on the same domain")
		return
	}

	// get normalized version of rawcurrenturl
	rcuNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error with normalize raw current url")
	}

	// if pages map already has rcuNormalized, increment count
	if _, ok := pages[rcuNormalized]; ok {
		pages[rcuNormalized] += 1
		return
	} else {
		// add the key and set it to 1
		pages[rcuNormalized] = 1
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
	urls, err := getURLsFromHTML(currentHTML, rawBaseURL)
	// fmt.Println("getting urls from html", urls)
	if err != nil {
		fmt.Println("error getting urls:", err)
		return
	}

	// crawl over every url in urls
	for _, u := range urls {
		crawlPage(rawBaseURL, u, pages)
	}
}
