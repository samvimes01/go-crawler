package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("skipping broken base URL %s\n", rawBaseURL)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("skipping broken current URL %s\n", rawCurrentURL)
		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("skipping external URL %s\n", rawCurrentURL)
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("skipping broken normalized URL %s\n", rawCurrentURL)
		return
	}

	if _, visited := pages[normalizedURL]; visited {
		fmt.Printf("skipping duplicate URL %s\n", rawCurrentURL)
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getHTML %s: %s\n", rawCurrentURL, err)
		return
	}

	fmt.Printf("Recursively crawling %s\n", normalizedURL)

	urls, err := getURLsFromHTML(html, currentURL)
	if err != nil {
		fmt.Printf("error getURLsFromHTML %s: %s\n", normalizedURL, err)
		return
	}
	fmt.Printf("Found %d URLs on %s\n", len(urls), normalizedURL)

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
