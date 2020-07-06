package main

// WebsiteChecker used for check a url whether it is validate
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites used for check urls validate
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// results[u] = wc(u)
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
