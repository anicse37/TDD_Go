package concurrency

type WebsiteChecker func(string) bool

type ChannelResult struct {
	string
	bool
}

func Checker(wc WebsiteChecker, websiteURL []string) map[string]bool {
	result := map[string]bool{}
	resultChannel := make(chan ChannelResult)
	for _, url := range websiteURL {
		go func() {
			resultChannel <- ChannelResult{url, wc(url)}
		}()
	}
	for i := 0; i < len(websiteURL); i++ {
		temp := <-resultChannel
		result[temp.string] = temp.bool
	}
	return result
}
