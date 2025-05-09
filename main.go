package main

import "bet-engine/bootstrap"

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"sync"

// 	"github.com/gorilla/mux"
// )

// var (
// 	bets     = make(map[int][]*Bet)  // userID -> list of bets
// 	balances = make(map[int]float64) // userID -> balance
// 	mu       = sync.RWMutex{}
// )

// func placeBet(w http.ResponseWriter, r *http.Request) {
// 	var bet Bet
// 	if err := json.NewDecoder(r.Body).Decode(&bet); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		log.Println("Error decoding bet:", err)
// 		return
// 	}
// 	bet.Status = "pending"

// 	mu.Lock()
// 	defer mu.Unlock()
// 	bets[bet.UserID] = append(bets[bet.UserID], &bet)
// 	balances[bet.UserID] -= bet.Amount
// 	log.Printf("Placed bet: %+v\n", bet)

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(bet)
// }

// func settleBet(w http.ResponseWriter, r *http.Request) {
// 	var req struct {
// 		EventID int    `json:"event_id"`
// 		Result  string `json:"result"` // win or lose
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		log.Println("Error decoding settle request:", err)
// 		return
// 	}

// 	mu.Lock()
// 	defer mu.Unlock()

// 	for userID, userBets := range bets {
// 		for _, bet := range userBets {
// 			if bet.EventID == req.EventID && bet.Status == "pending" {
// 				if req.Result == "win" {
// 					winAmount := bet.Amount * bet.Odds
// 					balances[userID] += winAmount
// 					bet.Status = "win"
// 					log.Printf("User %d won %.2f on event %d\n", userID, winAmount, req.EventID)
// 				} else if req.Result == "lose" {
// 					bet.Status = "lose"
// 					log.Printf("User %d lost bet on event %d\n", userID, req.EventID)
// 				}
// 			}
// 		}
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Event settled"))
// }

// func getBalance(w http.ResponseWriter, r *http.Request) {
// 	userIDStr := mux.Vars(r)["userID"]
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	mu.RLock()
// 	defer mu.RUnlock()

// 	balance := balances[userID]
// 	json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
// }

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/bets", placeBet).Methods("POST")
	// r.HandleFunc("/settle", settleBet).Methods("POST")
	// r.HandleFunc("/balance/{userID}", getBalance).Methods("GET")

	// log.Println("Starting server at :8080")
	// log.Fatal(http.ListenAndServe(":8080", r))

	bootstrap.Init()
}
