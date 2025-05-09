package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBalance(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   map[string]float64
	}{
		{
			name:           "Valid user ID",
			userID:         "1",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]float64{"balance": 100.0},
		},
		{
			name:           "Invalid user ID",
			userID:         "invalid",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	Balances = map[int]float64{
		1: 100.0,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/balance/"+tt.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			GetBalance(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedBody != nil {
				var responseBody map[string]float64
				if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
					t.Fatal(err)
				}
				if responseBody["balance"] != tt.expectedBody["balance"] {
					t.Errorf("expected balance %v, got %v", tt.expectedBody["balance"], responseBody["balance"])
				}
			}
		})
	}
}
