package router

import (
	"bet-engine/http/router/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/bets", controllers.PlaceBet).Methods("POST")
	r.HandleFunc("/settle", controllers.SettleBet).Methods("POST")
	r.HandleFunc("/balance/{userID}", controllers.GetBalance).Methods("GET")

	return r
}
