package selecturl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	selecturl "github.com/anicse37/TDD_Go/Select"
)

func TestSpeed(t *testing.T) {
	t.Run("Normal Test", func(t *testing.T) {

		url1 := "www.youtube.com"
		url2 := "www.google.com"
		got := selecturl.Racer(url1, url2)
		want := url2
		if got != want {
			t.Errorf("Got %v || Want %v\n", got, want)
		}
	})

	t.Run("Within 10 sec", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got := selecturl.Racer2(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("More than 10 sec", func(t *testing.T) {
		/*To see the program in action*/
		/*Change time below from 1 and 2 seconds to 12 and 11 seconds respectively*/
		slowServer := makeDelayedServer(2 * time.Second)
		fastServer := makeDelayedServer(1 * time.Second)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := selecturl.Racer3(slowURL, fastURL)
		if err != nil {
			t.Fatalf("Time Out")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

/*-------------------------------------------------------------------------------------------------*/

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
