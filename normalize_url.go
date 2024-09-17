package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {

	parseURL, err := url.Parse(u)
	if err != nil {
		fmt.Println("error parsing the url for normalization")
		return "", err
	}

	normalizedURL := parseURL.Host + parseURL.Path
	normalizedURL = strings.ToLower(normalizedURL)

	// trim trailing "/" from path
	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	// if the url is relative, build up absolute

	return normalizedURL, nil
}
