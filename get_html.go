package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		fmt.Printf("Error creating request %s: %s\n", rawURL, err)
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", rawURL, err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("status code error: %d - %s", resp.StatusCode, resp.Status)
	}
	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "text/html") {
		return "", fmt.Errorf("content type error: %s", ct)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body %s: %s\n", rawURL, err)
		return "", err
	}
	return string(body), nil
}
