package main

import (
	"log"
	"net/http"
	"os"
)

func coffeeHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()
	// Hata mesajı dönmesi..
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Hostname bazlı kahve sipariş başarıyla alındı. 
	w.Write([]byte("Kahve siparisiniz alinmistir. Talep eden kisi : " + servant))
}

func main() {
	http.HandleFunc("/coffee", coffeeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
