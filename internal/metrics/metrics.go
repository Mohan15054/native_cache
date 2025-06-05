package metrics

import "sync"

type Metrics interface {
	IncHits()
	IncMisses()
	Report() map[string]int
}

type BasicMetrics struct {
	mu     sync.Mutex
	hits   int
	misses int
}

func NewBasicMetrics() *BasicMetrics {
	return &BasicMetrics{}
}

func (m *BasicMetrics) IncHits() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hits++
}

func (m *BasicMetrics) IncMisses() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.misses++
}

func (m *BasicMetrics) Report() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return map[string]int{
		"hits":   m.hits,
		"misses": m.misses,
	}
}
