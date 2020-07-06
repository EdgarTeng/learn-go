package main

// WebsiteChecker used for check a url whether it is validate
type WebsiteChecker func(string) bool

// CheckWebsites used for check urls validate
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}
