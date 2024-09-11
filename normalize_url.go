package main

import (
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {

	normalizedURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	// trim trailing "/" from path
	normalizedURL.Path = strings.TrimSuffix(normalizedURL.Path, "/")

	// build up the desired url: blog.boot.dev/path
	return normalizedURL.Host + normalizedURL.Path, nil

}
