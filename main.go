package main

import (
	"encoding/json"
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
	decoder := json.NewDecoder(r.Body)
	var e Event
	err := decoder.Decode(&e)
	if err != nil {
		panic(err)
	}
	switch e.Type {
	case Deposit:
		if acc, _ := GetAccount(e.Destination); acc != nil {
			CreateNewDepositEvent(e, acc)
			eventResponse := EventDepositResponse{Destination: *acc}
			response, _ := json.Marshal(eventResponse)
			w.WriteHeader(http.StatusCreated)
			w.Write(response)
		} else {
			newAcc := Account{e.Destination, e.Amount}
			created := CreateNewAccount(newAcc)
			eventResponse := EventDepositResponse{Destination: *created}
			response, _ := json.Marshal(eventResponse)
			w.WriteHeader(http.StatusCreated)
			w.Write(response)
		}
	case Withdraw:
		if acc, _ := GetAccount(e.Origin); acc != nil {
			CreateNewWithdrawEvent(e, acc)
			eventResponse := EventWithdrawResponse{Origin: *acc}
			response, _ := json.Marshal(eventResponse)
			w.WriteHeader(http.StatusCreated)
			w.Write(response)
		} else {
			http.Error(w, "0", http.StatusNotFound)
		}
	case Transfer:
		originAcc, _ := GetAccount(e.Origin)
		destinationAcc, _ := GetAccount(e.Destination)
		if originAcc != nil {
			if destinationAcc == nil {
				newAcc := Account{e.Destination, 0}
				created := CreateNewAccount(newAcc)
				CreateNewTransferEvent(e, originAcc, created)
				eventResponse := EventTransferResponse{Origin: *originAcc, Destination: *created}
				response, _ := json.Marshal(eventResponse)
				w.WriteHeader(http.StatusCreated)
				w.Write(response)
			} else {
				CreateNewTransferEvent(e, originAcc, destinationAcc)
				eventResponse := EventTransferResponse{Origin: *originAcc, Destination: *destinationAcc}
				response, _ := json.Marshal(eventResponse)
				w.WriteHeader(http.StatusCreated)
				w.Write(response)
			}
		} else {
			http.Error(w, "0", http.StatusNotFound)
		}
	}
}

func handleGetAllAccounts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v", currentAccounts)))
}

func main() {
	fmt.Println("Building an API for EBANX")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /reset", handleReset)
	mux.HandleFunc("GET /balance", handleGetAccountBalance)
	mux.HandleFunc("GET /accounts", handleGetAllAccounts)
	mux.HandleFunc("POST /event", handlePostNewEvent)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}

}
