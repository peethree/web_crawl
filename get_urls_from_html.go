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
		return []string{}, err
	}

	urls := []string{}

	//  recursively traverse the node tree and find the <a> tag "anchor" elements
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// if path is relative, add a.Val ontop of input URL
					// else add new url to return list
					val := a.Val
					parsedUrl, err := url.Parse(val)
					if err != nil {
						continue
					}
					// url is relative
					if !parsedUrl.IsAbs() {
						val = rawBaseURL + parsedUrl.Path
					}
					urls = append(urls, val)
					break
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
