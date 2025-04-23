package contextProgram_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	contextProgram "github.com/anicse37/TDD_Go/Context"
)

type MockStore struct {
	response  string
	cancelled bool
}

func (m *MockStore) Fetch() string {
	time.Sleep(100 * time.Microsecond)
	return m.response
}
func (m *MockStore) Cancel() {
	m.cancelled = true
}
func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &MockStore{response: data}
		svr := contextProgram.Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &MockStore{response: data}
		svr := contextProgram.Server2(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		/*
		*Comment the cancel and un-comment the time.AfterFunc to see the program in action
		 */
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}
