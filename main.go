package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v\n", err)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - maxPages: %v\n", err)
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	// fmt.Printf("\n\n======\n\nPages found: %d\n", len(cfg.pages))
	// cnt := 1
	// for normalizedURL, _ := range cfg.pages {
	// 	fmt.Printf("%d - %s\n", cnt, normalizedURL)
	// 	cnt++
	// }
	if err := writeCSVReport(cfg.pages, "report.csv"); err != nil {
		fmt.Printf("Error - writeCSVReport: %v\n", err)
		return
	}
}
