package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	))

	router.Handle("/players/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			playerName := strings.TrimPrefix(r.URL.Path, "/players/")
			switch r.Method {
			case http.MethodPost:
				p.processWin(w, playerName)
			case http.MethodGet:
				p.showScore(w, playerName)
			}
		},
	))

	router.ServeHTTP(w,r)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, playerName string) {
	score := p.store.GetPlayerScore(playerName)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, playerName string) {
	p.store.RecordWin(playerName)
	w.WriteHeader(http.StatusAccepted)
}
