package contextProgram_test

import (
	contextProgram "aniket/Context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
}
