package concurrency_test

import (
	"reflect"
	"testing"
	"time"

	concurrency "github.com/anicse37/TDD_Go/Concurrency"
)

func MockWebsiteChecker(website string) bool {
	return !(website == "www.github.com")
}

func TestWebsiteTester(t *testing.T) {
	websites := []string{
		"www.google.com",
		"www.youtube.com",
		"www.github.com",
	}
	want := map[string]bool{
		"www.google.com":  true,
		"www.youtube.com": true,
		"www.github.com":  false,
	}
	got := concurrency.Checker(MockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v || Want %v \n", got, want)
	}

}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkTest(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "www.somewebsite.com"
	}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		concurrency.Checker(slowStubWebsiteChecker, urls)
	}
}
