package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type playersWrap struct {
	Players []player `json:"players"`
}

type players struct {
	sync.Mutex
	data playersWrap
}

var (
	ps = players{
		data: playersWrap{
			Players: []player{
				{ID: 1, Name: "Sally", Level: 2},
				{ID: 2, Name: "Lance", Level: 1},
				{ID: 3, Name: "Aki", Level: 3},
				{ID: 4, Name: "Maria", Level: 4},
				{ID: 5, Name: "Julio", Level: 1},
				{ID: 6, Name: "Julian", Level: 1},
				{ID: 7, Name: "Jaime", Level: 1},
			},
		},
	}
)

func playersHandler(w http.ResponseWriter, r *http.Request) {
	ps.Lock()
	defer ps.Unlock()

	if err := json.NewEncoder(w).Encode(ps.data); err != nil {
		stts := http.StatusInternalServerError
		http.Error(w, http.StatusText(stts), stts)
		return
	}
}

func main() {
	if err := http.ListenAndServe(":4000", http.HandlerFunc(playersHandler)); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
