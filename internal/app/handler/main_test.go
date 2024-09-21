package handler

import (
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

func TestMainPage(t *testing.T) {

	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "POST request",
			method:         http.MethodPost,
			body:           "https://example.com",
			expectedStatus: http.StatusCreated,
			expectedBody:   "http://localhost:8080/",
		},
		{
			name:           "Invalid method",
			method:         http.MethodPut,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/", strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(MainPage)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if tt.method == http.MethodPost {
				if !strings.HasPrefix(rr.Body.String(), tt.expectedBody) {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), tt.expectedBody)
				}
			} else if tt.method == http.MethodGet {
				if location := rr.Header().Get("Location"); location != "https://example.com" {
					t.Errorf("handler returned unexpected location header: got %v want %v",
						location, "https://example.com")
				}
			} else {
				if rr.Body.String() != tt.expectedBody {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), tt.expectedBody)
				}
			}
		})
	}
}
