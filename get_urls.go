package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getAttr(attr string, baseURL *url.URL, urls *[]string) func(i int, s *goquery.Selection) {
	return func(i int, s *goquery.Selection) {
		if val, exists := s.Attr(attr); exists && strings.TrimSpace(val) != "" {
			u, err := url.Parse(val)
			if err != nil {
				log.Printf("couldn't parse %v %q: %v\n", attr, val, err)

				return
			}

			absolute := baseURL.ResolveReference(u)
			*urls = append(*urls, absolute.String())
		}
	}
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	urls := []string{}

	doc.Find("a[href]").Each(getAttr("href", baseURL, &urls))

	if len(urls) == 0 {
		return nil, nil
	}

	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	urls := []string{}

	doc.Find("img[src]").Each(getAttr("src", baseURL, &urls))

	return urls, nil
}
