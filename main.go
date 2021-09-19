package main

import (
	"log"
	"net/http"
)

const (
	Port        = 7000
	checkoutUrl = "https://paypage.ru/api/url/"
	merchantID  = "133264"
	key         = "Testterminal1"
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
}

/////          CHARGE
////  		   ДОПЫ К CHARGE
////           BLOCK dealid
