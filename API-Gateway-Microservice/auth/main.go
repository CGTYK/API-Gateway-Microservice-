package main

import (
	"log"
	"net/http"
)

func checkAuth(w http.ResponseWriter, r *http.Request) {
	authString := r.Header.Get("Authorization")
	// Yetki doğruması başarılı..
	if authString == "CSlkjdfj3423lkj234jj==" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized: true"))
		return
	}
	// Yetkisi olmaığında.. 
	http.Error(w, "Authorized: false", http.StatusUnauthorized)
}

func main() {
	http.HandleFunc("/authenticated", checkAuth)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
