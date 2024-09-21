package handler

import (
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
