package controllers

import (
	"bet-engine/domain/entities"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const (
	WIN     = 0
	LOSE    = 1
	PENDING = 2
)

func SettleBet(w http.ResponseWriter, r *http.Request) {
	req := entities.SettleReq{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding settle request:", err)
		return
	}
	if req.Result != WIN && req.Result != LOSE {
		http.Error(w, "Result must be 'win' or 'lose'", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	userBetsCopy := make(map[int][]*entities.Bet)

	MU.RLock()
	for userID, userBets := range Bets {
		userBetsCopy[userID] = append([]*entities.Bet{}, userBets...)
	}
	MU.RUnlock()

	for userID, userBets := range userBetsCopy {
		wg.Add(1)
		go func(uid int, bets []*entities.Bet) {
			defer wg.Done()

			updated := false
			var winnings float64

			// Check if the user has any pending bets for the event
			hasPendingBets := false

			for _, bet := range bets {
				if bet.EventID == req.EventID && bet.Status == PENDING {
					hasPendingBets = true
					MU.Lock()
					if req.Result == WIN {
						winAmount := bet.Amount * bet.Odds
						Balances[uid] += winAmount
						winnings += winAmount
						bet.Status = WIN
					} else {
						bet.Status = LOSE
					}
					updated = true
					MU.Unlock()
				}
			}

			// If no pending bets found for the event, log this information
			if !hasPendingBets {
				log.Printf("No pending bets for user %d on event %d\n", uid, req.EventID)
			}

			if updated {
				log.Printf("Settled bets for user %d on event %d (Result: %d, Winnings: %.2f)\n", uid, req.EventID, req.Result, winnings)
			}
		}(userID, userBets)
	}

	wg.Wait()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event settled"))
}
