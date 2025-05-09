package controllers

import (
	"bet-engine/domain/entities"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := entities.UserCreateReq{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding user creation request:", err)
		return
	}

	if req.UserID <= 0 || req.Balance < 0 {
		http.Error(w, "User ID must be > 0 and balance must be >= 0", http.StatusBadRequest)
		return
	}

	MU.Lock()
	defer MU.Unlock()

	if _, exists := Balances[req.UserID]; exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	Balances[req.UserID] = req.Balance
	Bets[req.UserID] = []*entities.Bet{}
	log.Printf("Created user %d with balance %.2f\n", req.UserID, req.Balance)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": req.UserID,
		"balance": req.Balance,
	})
}
