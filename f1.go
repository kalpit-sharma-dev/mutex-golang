package main

import (
	"fmt"
	"sync"
)

type SafeStringList struct {
	mu   sync.Mutex
	data []string
}

// Add appends a string to the list in a thread-safe manner.
func (s *SafeStringList) Add(item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, item)
}

// Replace overrides the entire list with a new one in a thread-safe manner.
func (s *SafeStringList) Replace(newList []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append([]string{}, newList...) // Replace with a copy of the new list
}

// Get retrieves the list in a thread-safe manner.
func (s *SafeStringList) Get() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string{}, s.data...) // Return a copy to avoid data races
}

func main() {
	sharedList := &SafeStringList{}
	wg := sync.WaitGroup{}

	// Add items concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sharedList.Add(fmt.Sprintf("Item %d", id))
		}(i)
	}

	// Wait for additions to complete
	wg.Wait()

	fmt.Println("Before Replace:", sharedList.Get())

	// Replace the list
	sharedList.Replace([]string{"NewItem 1", "NewItem 2", "NewItem 3"})

	fmt.Println("After Replace:", sharedList.Get())
}
