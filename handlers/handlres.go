package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/store"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

type URLHandler struct {
	Store *store.URLStore
}

func NewURLHandler(store *store.URLStore) *URLHandler {
	return &URLHandler{Store: store}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	originalURL := req.URL

	if shortURL, exists := h.Store.GetShortURL(originalURL); exists {
		response := ShortenResponse{ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortURL)}
		json.NewEncoder(w).Encode(response)
		return
	}

	shortURL := store.GenerateShortURL()
	h.Store.SetURLMapping(shortURL, originalURL)

	response := ShortenResponse{ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortURL)}
	json.NewEncoder(w).Encode(response)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	if originalURL, exists := h.Store.GetOriginalURL(shortURL); exists {
		http.Redirect(w, r, originalURL, http.StatusFound)
		return
	}
	http.NotFound(w, r)
}

func (h *URLHandler) TopDomains(w http.ResponseWriter, r *http.Request) {
	topDomains := h.Store.GetTopDomains()
	json.NewEncoder(w).Encode(topDomains)
}
