package controllers

import (
	"bet-engine/domain/entities"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var (
	Bets     = make(map[int][]*entities.Bet)
	Balances = make(map[int]float64)
	MU       = sync.RWMutex{}
)

// func PlaceBet(w http.ResponseWriter, r *http.Request) {
// 	var bet entities.Bet
// 	if err := json.NewDecoder(r.Body).Decode(&bet); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		log.Println("Error decoding bet:", err)
// 		return
// 	}
// 	bet.Status = "pending"

// 	MU.Lock()
// 	defer MU.Unlock()
	
// 	Bets[bet.UserID] = append(Bets[bet.UserID], &bet)
// 	Balances[bet.UserID] -= bet.Amount
// 	log.Printf("Placed bet: %+v\n", bet)

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(bet)
// }

func PlaceBet(w http.ResponseWriter, r *http.Request) {
	var bet entities.Bet
	if err := json.NewDecoder(r.Body).Decode(&bet); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding bet:", err)
		return
	}
	if bet.Amount <= 0 || bet.Odds <= 1.0 {
		http.Error(w, "Invalid bet amount or odds", http.StatusBadRequest)
		return
	}

	MU.Lock()
	defer MU.Unlock()

	if Balances[bet.UserID] < bet.Amount {
		http.Error(w, "Insufficient balance", http.StatusBadRequest)
		return
	}

	bet.Status = "pending"
	Bets[bet.UserID] = append(Bets[bet.UserID], &bet)
	Balances[bet.UserID] -= bet.Amount
	log.Printf("Placed bet: %+v\n", bet)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bet)
}