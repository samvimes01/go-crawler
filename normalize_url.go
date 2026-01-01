package main

import (
	"fmt"
	"log"
	"strings"
	"net/url"
)

func normalizeURL(rawUrl string) (string, error) {
	trimmed := strings.TrimRight(rawUrl, "/")
	u, err := url.Parse(trimmed)
	if err != nil {
		log.Print(err)
		return "", err
	}

	return fmt.Sprintf("%s%s", u.Host, u.Path), nil
}
