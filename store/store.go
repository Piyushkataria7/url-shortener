package store

import (
	"math/rand"
	"sync"
	"time"
)

type URLStore struct {
	sync.RWMutex
	urlMap map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		urlMap: make(map[string]string),
	}
}

func (s *URLStore) GetOriginalURL(shortURL string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	originalURL, exists := s.urlMap[shortURL]
	return originalURL, exists
}

func (s *URLStore) SetURLMapping(shortURL, originalURL string) {
	s.Lock()
	defer s.Unlock()
	s.urlMap[shortURL] = originalURL
}

func GenerateShortURL() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}
