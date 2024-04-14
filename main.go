package main

import (
	"fmt"
	"net/http"
)

func handleReset(w http.ResponseWriter, r *http.Request) {
	ResetDB()
	_, err := w.Write([]byte("OK"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handleGetAccountBalance(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("account_id"); id == "" {
		http.Error(w, "0", http.StatusBadRequest)
	} else {
		accountX, err := GetAccount(id)
		if err == nil {
			w.Write([]byte(fmt.Sprintf("%v", accountX.Balance)))
		} else {
			http.Error(w, "0", http.StatusNotFound)
		}
	}
}

func handlePostNewEvent(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	fmt.Println(body)
}

func main() {
	fmt.Println("Building an API for EBANX")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /reset", handleReset)
	mux.HandleFunc("GET /balance", handleGetAccountBalance)
	mux.HandleFunc("POST /event", handlePostNewEvent)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}

}
