package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a website")
		os.Exit(1)
	}

	website := os.Args[1:][0]

	if len(website) == 0 {
		fmt.Println("Please provide a website")
		os.Exit(1)
	}

	visited := make(map[string]bool)

	scrape(website, visited)
}

func scrape(page string, visited map[string]bool) {
	if visited[page] {
		return
	}

	visited[page] = true

	doc, err := loadPage(page)

	if err != nil {
		fmt.Println(err)
		return
	}

	if doc == nil {
		return
	}

	hrefs := getAnchors(doc)

	for _, href := range hrefs {
		if strings.HasPrefix(href, "/") {
			href = page + href
			scrape(href, visited)
		}

		if strings.HasPrefix(href, page) {
			scrape(href, visited)
		}

		checkLinks(href, visited)
	}
}

func loadPage(url string) (*html.Node, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode >= 400 {
		fmt.Println(url + " is not reachable")
		return nil, nil
	} else {
		fmt.Println(url + " is reachable")
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func getAnchors(n *html.Node) []string {
	var hrefs []string
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					hrefs = append(hrefs, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(n)
	return hrefs
}

func checkLinks(page string, visited map[string]bool) {
	if visited[page] {
		return
	}

	visited[page] = true

	data, err := http.Get(page)

	if err != nil {
		fmt.Println(page + " is not reachable")
		return
	}

	if data.StatusCode >= 400 {
		fmt.Println(page + " is not reachable")
		return
	}

	fmt.Println(page + " is reachable")
}
