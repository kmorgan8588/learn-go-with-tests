package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
func TestRacer(t *testing.T) {
	t.Run("fast vs slow", func(t *testing.T) {

		slowServer := makeTestServer(20 * time.Millisecond)
		fastServer := makeTestServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an erro but got one e%v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond with a specified time", func(t *testing.T) {
		server := makeTestServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expect an error but didn't get one")
		}
	})
}
