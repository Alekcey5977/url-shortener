package storage

import (
	"sync"
	"url-shortener/internal/models"
)

type Storage interface {
	Save(urlMapping *models.URLMapping) error
	Find(shortURL string) (*models.URLMapping, error)
}

type MemoryStorage struct {
	urls map[string]*models.URLMapping
	mu sync.RWMutex
}

func NewMemorySrorage() *MemoryStorage {
	return &MemoryStorage{
		urls: make(map[string]*models.URLMapping),
	}
}

// Save - сохраняет URL в памяти
func (s *MemoryStorage) Save(urlMapping *models.URLMapping) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.urls[urlMapping.ShortURL] = urlMapping
	return nil
}

// Find - ищет URL по короткой ссылке
func (s *MemoryStorage) Find(shortURL string) (*models.URLMapping, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    if url, exists := s.urls[shortURL]; exists {
        return url, nil
    }
    
    return nil, nil // Вернем nil если не найдено
}