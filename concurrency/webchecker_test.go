package main

import (
	"reflect"
	"testing"
	"time"
)

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://youtube.com",
		"https://wikipedia.com",
		"https://amazon.com",
		"notgood://1234.ddf",
	}

	actual := CheckWebsites(mockWebSiteChecker, websites)

	expect := map[string]bool{
		"https://google.com":    true,
		"https://youtube.com":   true,
		"https://wikipedia.com": true,
		"https://amazon.com":    true,
		"notgood://1234.ddf":    false,
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Fatalf("expect %v, actual %v", expect, actual)
	}

}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebSiteChecker, urls)
	}
}

func mockWebSiteChecker(url string) bool {
	if url == "notgood://1234.ddf" {
		return false
	}
	return true
}

func slowWebSiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
