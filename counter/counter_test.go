package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	reset()

	if val := get(); val != 0 {
		t.Errorf("Expected counter to be 0 after initialization, got %d", val)
	}

	if val := up(); val != 1 {
		t.Errorf("Expected counter to be 1, got %d", val)
	}
	if val := up(); val != 2 {
		t.Errorf("Expected counter to be 2, got %d", val)
	}

	if val := get(); val != 2 {
		t.Errorf("Expected counter to be 2, got %d", val)
	}

	reset()
	if val := get(); val != 0 {
		t.Errorf("Expected counter to be 0 after reset, got %d", val)
	}
}

func TestCounterThreadSafety(t *testing.T) {
	reset()

	const numVUs = 100
	const incrementsPerVU = 10
	var wg sync.WaitGroup

	wg.Add(numVUs)
	for i := 0; i < numVUs; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerVU; j++ {
				Increment()
			}
		}()
	}

	wg.Wait()

	expectedValue := numVUs * incrementsPerVU
	actualValue := get()

	if actualValue != expectedValue {
		t.Errorf("Expected counter to be %d, got %d", expectedValue, actualValue)
	}
}
