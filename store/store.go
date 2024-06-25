package store

import (
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type domainCount struct {
	Domain string
	Count  int
}
type URLStore struct {
	sync.RWMutex
	urlMap     map[string]string
	domainMap  map[string]int
	reverseMap map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		urlMap:     make(map[string]string),
		domainMap:  make(map[string]int),
		reverseMap: make(map[string]string),
	}
}

func (s *URLStore) GetOriginalURL(shortURL string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	originalURL, exists := s.urlMap[shortURL]
	return originalURL, exists
}

func (s *URLStore) SetURLMapping(shortURL, originalURL string) string {
	s.Lock()
	defer s.Unlock()

	if existingShortURL, exists := s.reverseMap[originalURL]; exists {
		return existingShortURL
	}

	s.urlMap[shortURL] = originalURL
	domain := ExtractDomain(originalURL)
	s.domainMap[domain]++

	return shortURL
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

func ExtractDomain(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return strings.TrimPrefix(parsedURL.Host, "www.")
}

func (s *URLStore) GetTopDomains() []string {
	s.RLock()
	defer s.RUnlock()

	var domains []domainCount
	for domain, count := range s.domainMap {
		domains = append(domains, domainCount{Domain: domain, Count: count})
	}

	sort.Slice(domains, func(i, j int) bool {
		return domains[i].Count > domains[j].Count
	})

	var topDomains []string
	for i := 0; i < len(domains) && i < 3; i++ {
		topDomains = append(topDomains, domains[i].Domain+": "+strconv.Itoa(domains[i].Count))
	}

	return topDomains
}

func (s *URLStore) GetShortURL(originalURL string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	for short, original := range s.urlMap {
		if original == originalURL {
			return short, true
		}
	}
	return "", false
}
