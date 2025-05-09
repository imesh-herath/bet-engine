package entities

type Bet struct {
	UserID  int     `json:"user_id"`
	EventID int     `json:"event_id"`
	Odds    float64 `json:"odds"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

type SettleReq struct {
	EventID int    `json:"event_id"`
	Result  string `json:"result"`
}

type UserCreateReq struct {
	UserID  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}
