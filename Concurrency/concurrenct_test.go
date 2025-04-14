package concurrency_test

import (
	concurrency "aniket/Concurrency"
	"reflect"
	"testing"
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
