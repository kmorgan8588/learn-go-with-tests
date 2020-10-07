package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

}
func TestRacer(t *testing.T) {
	t.Run("fast vs slow", func(t *testing.T) {

		slowServer := makeTestServer(20)
		fastServer := makeTestServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
