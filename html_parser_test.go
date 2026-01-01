package main

import (
	"testing"
)

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "return h1 content",
			inputURL: "<h1>content of h1</h1>",
			expected: "content of h1",
		},
		{
			name:     "return h1 content from html",
			inputURL: "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "returns an empty string if no <h1> tag is found.",
			inputURL: "<b>content of h1</b>",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputURL)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "return p content",
			inputURL: "<p>content of p</p>",
			expected: "content of p",
		},
		{
			name: "return p content from html",
			inputURL: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name:     "returns first <p> tag found.",
			inputURL: "<p>content of first p</p><p>content of second p</p>",
			expected: "content of first p",
		},
		{
			name:     "returns an empty string if no <p> tag is found.",
			inputURL: "<b>content of b</b>",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputURL)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
