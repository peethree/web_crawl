package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	// io.reader
	r := strings.NewReader(htmlBody)

	// html Node tree
	htmlNodeTree, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	// make a slice that holds all the urls
	urls := []string{}

	// parse base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	//  recursively traverse the html node tree and find the <a> tag "anchor" elements
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// if path is relative, reference absolute URL
					parsedUrl, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					// handle relative urls
					resolvedURL := baseURL.ResolveReference(parsedUrl)

					// append to return slice of urls as a string
					urls = append(urls, resolvedURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(htmlNodeTree)

	return urls, nil

}
