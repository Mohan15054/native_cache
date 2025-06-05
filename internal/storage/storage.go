package storage

import (
	"cachesystem/internal/config"
	"cachesystem/internal/eviction"
	"cachesystem/internal/metrics"
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiryTime time.Time
}

type InMemoryStore struct {
	mu      sync.RWMutex
	data    map[string]*item
	cfg     *config.CacheConfig
	evict   eviction.EvictionPolicy
	metrics metrics.Metrics
}

func NewInMemoryStore(cfg *config.CacheConfig, evict eviction.EvictionPolicy, metrics metrics.Metrics) *InMemoryStore {
	store := &InMemoryStore{
		data:    make(map[string]*item),
		cfg:     cfg,
		evict:   evict,
		metrics: metrics,
	}
	go store.cleanupExpired()
	return store
}

func (s *InMemoryStore) Set(key string, value interface{}, ttl int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	expiry := time.Now().Add(time.Duration(ttl) * time.Second)
	s.data[key] = &item{value: value, expiryTime: expiry}
	s.evict.OnInsert(key)
	return nil
}

func (s *InMemoryStore) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	it, ok := s.data[key]
	if !ok || time.Now().After(it.expiryTime) {
		s.metrics.IncMisses()
		return nil, false
	}
	s.evict.OnAccess(key)
	s.metrics.IncHits()
	return it.value, true
}

func (s *InMemoryStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return nil
}

func (s *InMemoryStore) cleanupExpired() {
	for {
		time.Sleep(1 * time.Minute)
		now := time.Now()
		s.mu.Lock()
		for k, v := range s.data {
			if now.After(v.expiryTime) {
				delete(s.data, k)
			}
		}
		s.mu.Unlock()
	}
}
