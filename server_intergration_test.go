package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	playerName := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(playerName))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
