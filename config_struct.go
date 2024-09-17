package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	// if the url is in the dictionary, that means it's not the first visit
	if _, ok := cfg.pages[normalizedURL]; ok {
		return false
	}
	// else add it to dict and return true
	cfg.pages[normalizedURL] = 1
	return true
}
