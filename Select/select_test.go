package selecturl_test

import (
	selecturl "aniket/Select"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSpeed(t *testing.T) {
	url1 := "www.youtube.com"
	url2 := "www.google.com"
	got := selecturl.Racer(url1, url2)
	want := url2
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}

func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := selecturl.Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	slowServer.Close()
	fastServer.Close()
}
