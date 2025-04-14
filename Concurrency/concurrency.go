package concurrency

import "fmt"

type WebsiteChecker func(string) bool

func Checker(wc WebsiteChecker, websiteURL []string) map[string]bool {
	result := map[string]bool{}
	for _, url := range websiteURL {
		result[url] = wc(url)
	}
	fmt.Println(result)
	return result

}
