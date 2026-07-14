package main

import (
	"fmt"
	"sync"
)

var (
	completionRegistered = make(map[string]bool)
	completionMu         sync.Mutex
)

func registerCompletion(name string) error {
	completionMu.Lock()
	defer completionMu.Unlock()
	
	if completionRegistered[name] {
		return fmt.Errorf("completion already registered for %s", name)
	}
	completionRegistered[name] = true
	return nil
}

func main() {
	err := registerCompletion("test")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Completion registered!")
	}
}
