package main

import (
	"log"
	"net/http"
	"os"
)

func teaHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()
	// Hata mesajı dönmesi..
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Hostname bazlı çay sipariş başarıyla alındı. 
	w.Write([]byte("Cay siparisiniz alinmistir. Talep eden kisi : " + servant))
}

func main() {
	http.HandleFunc("/tea", teaHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
