package contextProgram_test

import (
	contextProgram "aniket/Context"
	"context"
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
	t.Run("Cancel when Requested", func(t *testing.T) {
		data := "Hello World"
		mockStore := &MockStore{response: data}
		svr := contextProgram.Server(mockStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCTX, _ := context.WithCancel(request.Context())
		time.After(5 * time.Millisecond)
		request = request.WithContext(cancellingCTX)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if !mockStore.cancelled {
			t.Errorf("Store was not told to cancel")
		}
	})
}
