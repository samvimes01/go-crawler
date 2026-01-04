package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Println("No data to write to CSV")
		return nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}); err != nil {
		fmt.Printf("Error writing header to CSV: %v\n", err)
		return err
	}
	for pageURL, pageData := range pages {
		outgoingLinkURLs := strings.Join(pageData.OutgoingLinks, ",")
		imageURLs := strings.Join(pageData.ImageURLs, ",")

		if err := writer.Write([]string{pageURL, pageData.H1, pageData.FirstParagraph, outgoingLinkURLs, imageURLs}); err != nil {
			fmt.Printf("Error writing row to CSV: %v\n", err)
			return err
		}
	}

	return nil
}
