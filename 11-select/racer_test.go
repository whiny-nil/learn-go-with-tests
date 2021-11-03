package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns the url of the fastest site", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		expected := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatal("expected no error but got one")
		}

		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("it returns an error if neither site responds in 10s", func(t *testing.T) {
		server1 := makeDelayedServer(11 * time.Millisecond)
		server2 := makeDelayedServer(12 * time.Millisecond)

		defer server1.Close()
		defer server2.Close()

		_, err := ConfigurableRacer(server1.URL, server2.URL, (10 * time.Millisecond))

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
