package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type playersWrap struct {
	Players []player `json:"players"`
}

func playersHandler(w http.ResponseWriter, r *http.Request) {
	d := playersWrap{
		Players: []player{
			{ID: 1, Name: "Sally", Level: 2},
			{ID: 2, Name: "Lance", Level: 1},
			{ID: 3, Name: "Aki", Level: 3},
			{ID: 4, Name: "Maria", Level: 4},
		},
	}

	if err := json.NewEncoder(w).Encode(d); err != nil {
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
