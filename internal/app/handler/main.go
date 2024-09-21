package handler

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Agniy/shortener/internal/app/config"
	"io"
	"net/http"
	"strings"
)

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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	cfg := config.GetConfig()
	originalURL := string(body)
	shortURL := "http://" + cfg.App.IP + ":" + cfg.Port + "/" + generateShortURL(originalURL)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}

// redirect to original url
func redirectToOriginal(w http.ResponseWriter, r *http.Request) {
	// New GET logic
	id := strings.TrimPrefix(r.URL.Path, "/")
	if id != "" {
		originalURL := getOriginalURL(id)
		if originalURL != "" {
			w.Header().Set("Location", originalURL)
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
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
