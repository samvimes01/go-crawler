package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.maxPagesReached() {
		return
	}

	baseURL := cfg.baseURL
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("skipping broken current URL %s\n", rawCurrentURL)

		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("skipping broken normalized URL %s\n", rawCurrentURL)
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("Recursively crawling %s\n", normalizedURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getHTML %s: %s\n", rawCurrentURL, err)
		return
	}

	pageData := extractPageData(html, rawCurrentURL)
	cfg.setPageData(normalizedURL, pageData)

	for _, url := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
