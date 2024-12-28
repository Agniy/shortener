package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 8,
		},
		{
			name:     "Short URL",
			input:    "https://example.com",
			expected: 8,
		},
		{
			name:     "Long URL",
			input:    "https://www.example.com/very/long/path/with/many/segments",
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateShortURL(tc.input)
			if len(result) != tc.expected {
				t.Errorf("Expected length %d, got %d for input %s", tc.expected, len(result), tc.input)
			}
		})
	}
}

func TestShortenURL(t *testing.T) {
	tests := []struct {
		name           string
		requestURL     string
		expectedStatus int
		wantErr        bool
	}{
		{
			name:           "Valid URL",
			requestURL:     "https://practicum.yandex.ru",
			expectedStatus: http.StatusCreated,
			wantErr:        false,
		},
		{
			name:           "Empty URL",
			requestURL:     "",
			expectedStatus: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name:           "Invalid URL",
			requestURL:     "not-a-url",
			expectedStatus: http.StatusBadRequest,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			reqBody := ShortenRequest{
				URL: tt.requestURL,
			}
			bodyBytes, _ := json.Marshal(reqBody)

			// Create request
			req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Call handler
			MainPage(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("ShortenURL() status = %v, want %v", w.Code, tt.expectedStatus)
			}

			// For successful requests, verify response format
			if !tt.wantErr {
				var response ShortenResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}

				// Check if response contains a valid shortened URL
				if !strings.HasPrefix(response.Result, "http://") {
					t.Errorf("Response URL doesn't have correct prefix: %v", response.Result)
				}

				// Check if Content-Type header is set correctly
				contentType := w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Content-Type = %v, want application/json", contentType)
				}
			}
		})
	}
}
