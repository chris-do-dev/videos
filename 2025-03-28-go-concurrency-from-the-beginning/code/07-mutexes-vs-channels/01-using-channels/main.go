package main

import "context"

func inventoryManager(ctx context.Context, in <-chan func(map[string]int)) {
	inventory := map[string]int{}

	for {
		select {
		case <-ctx.Done():
			return
		case f := <-in:
			f(inventory)
		}
	}
}

type InventoryManager chan func(map[string]int)

func NewChannelScoreboardManager(ctx context.Context) InventoryManager {
	ch := make(InventoryManager)
	go inventoryManager(ctx, ch)
	return ch
}

func (cim InventoryManager) Update(name string, val int) {
	cim <- func(m map[string]int) {
		m[name] = val
	}
}

func (cim InventoryManager) Read(name string) (int, bool) {
	type Result struct {
		out int
		ok  bool
	}

	resultCh := make(chan Result)
	cim <- func(m map[string]int) {
		out, ok := m[name]
		resultCh <- Result{out, ok}
	}

	result := <-resultCh
	return result.out, result.ok
}
