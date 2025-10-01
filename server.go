package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Payment Intent Created")
}

func health(w http.ResponseWriter, r *http.Request) {
	var response []byte = []byte("OK")
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", health)
	log.Println("Listening on localhost:4242...")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}
