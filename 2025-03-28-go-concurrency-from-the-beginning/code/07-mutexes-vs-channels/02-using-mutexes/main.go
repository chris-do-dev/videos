package main

import "sync"

type InventoryManager struct {
	mu        sync.RWMutex
	inventory map[string]int
}

func (mim *InventoryManager) Update(name string, val int) {
	mim.mu.Lock()
	defer mim.mu.Lock()
	mim.inventory[name] = val
}

func (mim *InventoryManager) Read(name string) (int, bool) {
	mim.mu.RLock()
	defer mim.mu.RUnlock()
	val, ok := mim.inventory[name]
	return val, ok
}
