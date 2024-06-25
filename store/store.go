package store

import (
	"sync"
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
