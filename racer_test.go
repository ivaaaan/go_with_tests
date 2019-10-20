package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns faster url", func(t *testing.T) {
		slowServer := makeDelayedServer(5 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL, 30*time.Millisecond)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if err != nil {
			t.Errorf("got an error, but didn't expect one")
		}
	})

	t.Run("returns an error if server doesn't response within defined time", func(t *testing.T) {
		server := makeDelayedServer(10 * time.Millisecond)

		defer server.Close()

		_, err := Racer(server.URL, server.URL, 5*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
