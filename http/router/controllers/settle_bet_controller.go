package controllers

import (
	"bet-engine/domain/entities"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// func SettleBet(w http.ResponseWriter, r *http.Request) {
// 	req := entities.Req{}

// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		log.Println("Error decoding settle request:", err)
// 		return
// 	}

// 	MU.Lock()
// 	defer MU.Unlock()

// 	for userID, userBets := range Bets {
// 		for _, bet := range userBets {
// 			if bet.EventID == req.EventID && bet.Status == "pending" {
// 				if req.Result == "win" {
// 					winAmount := bet.Amount * bet.Odds
// 					Balances[userID] += winAmount
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

func SettleBet(w http.ResponseWriter, r *http.Request) {
	req := entities.SettleReq{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding settle request:", err)
		return
	}
	if req.Result != "win" && req.Result != "lose" {
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

			for _, bet := range bets {
				if bet.EventID == req.EventID && bet.Status == "pending" {
					MU.Lock()
					if req.Result == "win" {
						winAmount := bet.Amount * bet.Odds
						Balances[uid] += winAmount
						winnings += winAmount
						bet.Status = "win"
					} else {
						bet.Status = "lose"
					}
					updated = true
					MU.Unlock()
				}
			}

			if updated {
				log.Printf("Settled bets for user %d on event %d (Result: %s, Winnings: %.2f)\n", uid, req.EventID, req.Result, winnings)
			}
		}(userID, userBets)
	}

	wg.Wait()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event settled"))
}
