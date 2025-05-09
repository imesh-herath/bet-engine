package controllers

import (
	"bet-engine/domain/entities"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlaceBet(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        entities.Bet
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid bet",
			reqBody: entities.Bet{
				UserID: 1,
				Amount: 10.0,
				Odds:   2.0,
				Status: 0,
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"user_id": 1,
				"amount":  10.0,
				"odds":    2.0,
				"status":  0,
			},
		},
		{
			name: "Invalid user ID",
			reqBody: entities.Bet{
				UserID: 0,
				Amount: 10.0,
				Odds:   2.0,
				Status: 0,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	Balances = map[int]float64{
		1: 100.0,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.reqBody)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest("POST", "/place_bet", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()

			// Call CreateUser handler
			handler := http.HandlerFunc(CreateUser)
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedBody != nil {
				var responseBody map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
					t.Fatal(err)
				}
				if responseBody["user_id"] != tt.expectedBody["user_id"] ||
					responseBody["amount"] != tt.expectedBody["amount"] ||
					responseBody["odds"] != tt.expectedBody["odds"] {
					t.Errorf("expected response body %+v, got %+v", tt.expectedBody, responseBody)
				}
			}
			if tt.expectedStatus == http.StatusCreated {
				var bet entities.Bet
				if err := json.Unmarshal(w.Body.Bytes(), &bet); err != nil {
					t.Fatal(err)
				}
				if bet.UserID != tt.reqBody.UserID || bet.Amount != tt.reqBody.Amount ||
					bet.Odds != tt.reqBody.Odds || bet.Status != tt.reqBody.Status {
					t.Errorf("expected bet %+v, got %+v", tt.reqBody, bet)
				}
			}
			if tt.expectedStatus == http.StatusBadRequest {
				var errorResponse map[string]string
				if err := json.Unmarshal(w.Body.Bytes(), &errorResponse); err != nil {
					t.Fatal(err)
				}
				if errorResponse["error"] == "" {
					t.Errorf("expected error message, got empty")
				}
			}
		})
	}
}
