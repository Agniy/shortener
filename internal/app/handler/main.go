package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Agniy/shortener/internal/app/models"

	"net/http"
	"strings"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		makeShortenUrl(w, r)
	} else if r.Method == http.MethodGet {
		redirectToOriginal(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handle post request to get short url
func makeShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// check if url is valid
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(req.URL)
	baseURL := "http://" + r.Host + "/api/shorten"

	response := ShortenResponse{
		Result: baseURL + "/" + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// redirect to original url
func redirectToOriginal(w http.ResponseWriter, r *http.Request) {
	// New GET logic
	id := strings.TrimPrefix(r.URL.Path, "/api/shorten/")
	if id != "" {

		//try to get original url from db
		link, err := models.GetLink(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting link: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", link.URL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	} else {
		http.NotFound(w, r)
	}
}

func generateShortURL(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])[:8]
}

func getOriginalURL(id string) string {
	// This is a placeholder function. In a real implementation,
	// you would look up the original URL in your storage system.
	// For now, we'll just return a dummy URL.

	return "https://example.com"
}
