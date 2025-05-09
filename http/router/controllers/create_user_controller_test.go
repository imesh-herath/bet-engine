package controllers

import (
	"bet-engine/domain/entities"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        entities.UserCreateReq
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid user creation",
			reqBody: entities.UserCreateReq{
				UserID:  1,
				Balance: 100.0,
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"user_id": 1,
				"balance": 100.0,
			},
		},
		{
			name: "Invalid user ID",
			reqBody: entities.UserCreateReq{
				UserID:  0,
				Balance: 100.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Negative balance",
			reqBody: entities.UserCreateReq{
				UserID:  1,
				Balance: -10.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "User already exists",
			reqBody: entities.UserCreateReq{
				UserID:  1,
				Balance: 100.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	Balances = map[int]float64{
		1: 50.0,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request and recorder
			reqBody, _ := json.Marshal(tt.reqBody)
			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			// Call CreateUser handler
			handler := http.HandlerFunc(CreateUser)
			handler.ServeHTTP(rr, req)

			// Assert status code
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			// Assert response body
			var responseBody map[string]interface{}
			if err := json.NewDecoder(rr.Body).Decode(&responseBody); err != nil && tt.expectedStatus == http.StatusCreated {
				t.Fatal("failed to decode response body:", err)
			}

			if tt.expectedStatus == http.StatusCreated {
				// Compare expected and actual response body
				for key, expectedVal := range tt.expectedBody {
					if responseBody[key] != expectedVal {
						t.Errorf("expected %s: %v, got %v", key, expectedVal, responseBody[key])
					}
				}
			}
		})
	}
}
