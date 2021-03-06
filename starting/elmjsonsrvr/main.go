package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/codemodus/chain"
	"github.com/codemodus/mixmux"
	"github.com/codemodus/parth"
)

type player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type players struct {
	sync.Mutex
	data []*player
}

func (p *players) indexByID(id int) (int, bool) {
	for k, v := range p.data {
		if v.ID == id {
			return k, true
		}
	}

	return 0, false
}

var (
	ps = players{
		data: []*player{
			{ID: 1, Name: "Sally", Level: 2},
			{ID: 2, Name: "Lance", Level: 1},
			{ID: 3, Name: "Aki", Level: 3},
			{ID: 4, Name: "Maria", Level: 4},
			{ID: 5, Name: "Julio", Level: 1},
			{ID: 6, Name: "Julian", Level: 1},
			{ID: 7, Name: "Jaime", Level: 1},
		},
	}
)

func playersGetHandler(w http.ResponseWriter, r *http.Request) {
	ps.Lock()
	defer ps.Unlock()

	if err := json.NewEncoder(w).Encode(ps.data); err != nil {
		stts := http.StatusInternalServerError
		http.Error(w, http.StatusText(stts), stts)
		return
	}
}

func playersPatchHandler(w http.ResponseWriter, r *http.Request) {
	ps.Lock()
	defer ps.Unlock()

	id, err := parth.SegmentToInt(r.URL.Path, -1)
	if err != nil {
		stts := http.StatusBadRequest
		http.Error(w, http.StatusText(stts), stts)
		return
	}

	pi, ok := ps.indexByID(id)
	if !ok {
		stts := http.StatusNotFound
		http.Error(w, http.StatusText(stts), stts)
		return
	}

	p := ps.data[pi]
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		stts := http.StatusBadRequest
		http.Error(w, http.StatusText(stts), stts)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		stts := http.StatusInternalServerError
		http.Error(w, http.StatusText(stts), stts)
		return
	}
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := r.Header.Get("Origin")
		if o == "" {
			stts := http.StatusBadRequest
			http.Error(w, http.StatusText(stts), stts)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", o)

		next.ServeHTTP(w, r)
	})
}

func main() {
	c := chain.New(cors)
	m := mixmux.NewRouter(nil)

	m.Get("/players", c.EndFn(playersGetHandler))
	m.OptionsAuto("/players", c.End, nil)

	m.Patch("/players/:id", c.EndFn(playersPatchHandler))
	m.OptionsAuto("/players/:id", c.End, nil)

	if err := http.ListenAndServe(":4000", m); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
