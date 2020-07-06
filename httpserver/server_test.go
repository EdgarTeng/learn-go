package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	t.Run("return Ken's score", func(t *testing.T) {
		actual := response.Body.String()
		expect := "20"
		if actual != expect {
			t.Errorf("actual %s, but expect %s", actual, expect)
		}
	})
}
