package main

import (
	"testing"
)

// func TestNormalizeURL(t *testing.T) {
// 	got := normalizeURL("test")
// 	want := "stubbed"
// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove unsafe hypertext transfer protocol",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove extra '/' after path without scheme",
			inputURL: "blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove extra '/' after path with http scheme",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove extra '/' after path with https scheme",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "tests with domain",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "further tests with domain",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "query",
			inputURL: "https://blog.boot.dev/path?query=value",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "parameter",
			inputURL: "https://blog.boot.dev/path#section",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "subdomain",
			inputURL: "https://subdomain.blog.boot.dev/path",
			expected: "subdomain.blog.boot.dev/path",
		},
		{
			name:     "port",
			inputURL: "https://blog.boot.dev:8080/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "relative url",
			inputURL: "/path/to/page",
			expected: "/path/to/page",
		},
		{
			name:     "just a '/'",
			inputURL: "/",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
