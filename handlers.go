package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("get out and try again"))
		return
	}
	w.WriteHeader(http.StatusForbidden)
}

func block(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("get out, mfuc"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var br BlockRequest
	err := decoder.Decode(&br)
	if err != nil {
		fmt.Fprint(w, "request can't be decoded")
		return
	}

	card := br.Card

	if err := card.Verify(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	if err := card.Init(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	if err := WriteFile(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
}

///// рандом номера + deal для charge

func charge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("get out of here"))
		return
	}
}
