package main

import (
	"log"
	"net/http"
)

const (
	Port        = 7000
)

func main() {
	// change, err :=
	port := ":7000"
	http.HandleFunc("/", index)
	http.HandleFunc("/Block", block)
	http.HandleFunc("/Charge", charge)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
	fmt.Println("OK")
}

