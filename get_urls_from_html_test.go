package main

import (
	"reflect"
	"testing"
)

// func TestNormalizeURL(t *testing.T) {
// 	got := normalizeURL("test")
// 	want := "stubbed"
// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

func TestGetURLsfromHTML(t *testing.T) {

	// name := "absolute and relative URLs"
	inputURL := "https://blog.boot.dev"
	inputBody := `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="/okay/try/this/on/for/size">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`
	expected := []string{"https://blog.boot.dev/path/one", "https://other.com/path/one", "https://blog.boot.dev/okay/try/this/on/for/size"}

	// for i, tc := range tests {
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
	// 		if err != nil {
	// 			t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
	// 			return
	// 		}
	// 		if actual != tc.expected {
	// 			t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
	// 		}
	// 	})
	// }

	// 2 values, the list of urls and the error (here _)
	got, _ := getURLsFromHTML(inputBody, inputURL)

	// assert all <a> tags are found
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

}
