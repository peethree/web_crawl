package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// fetch webpage of rawURL
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	// if status code is 400+, return err
	if resp.StatusCode >= 400 {
		return "", errors.New("bad status code: " + http.StatusText(resp.StatusCode))
	}

	defer resp.Body.Close()

	// return error if content type header is not text/html
	contentType := resp.Header.Get("Content-Type")
	if contentType != "" && !strings.HasPrefix(contentType, "text/html") {
		return "", errors.New("content type is not text/html")
	}

	// return the html
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
