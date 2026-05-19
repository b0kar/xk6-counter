package counter

import (
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

var globalCounter = &Counter{}

func up() int {
	globalCounter.mu.Lock()
	defer globalCounter.mu.Unlock()
	globalCounter.value++
	return globalCounter.value
}

func get() int {
	globalCounter.mu.Lock()
	defer globalCounter.mu.Unlock()
	return globalCounter.value
}

func reset() {
	globalCounter.mu.Lock()
	defer globalCounter.mu.Unlock()
	globalCounter.value = 0
}
